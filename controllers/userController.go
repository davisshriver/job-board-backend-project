package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/davisshriver/job-board-backend-project/database"
	helper "github.com/davisshriver/job-board-backend-project/helpers"
	"github.com/davisshriver/job-board-backend-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db = database.GetDB()
var validate = validator.New()

type loginResponse struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("Login credentials are incorrect!")
		check = false
	}

	return check, msg
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// Check if the email or phone number already exists in the database
		var existingUser models.User
		err = db.Where("email = ?", *user.Email).Or("phone = ?", *user.Phone).First(&existingUser).Error
		if err == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "The email or phone number is already being used."})
			return
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while checking for user's email or phone number!"})
			return
		}

		// Generate tokens
		userId := generateUniqueUserId()

		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		user.UID = userId

		// Hash the password
		password := HashPassword(*user.Password)
		user.Password = &password

		// Create the user in the database using GORM
		err = db.Create(&user).Error
		if err != nil {
			msg := fmt.Sprintf("User item was not created properly!")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var foundUser models.User

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = db.Where("email = ?", user.Email).First(&foundUser).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Email or password is incorrect!"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		// Generate tokens
		token, refreshToken, err := helper.GenerateAllTokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, *foundUser.UserType)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while generating tokens!"})
			return
		}

		// Upsert the user tokens in the user_tokens table
		err = db.FirstOrCreate(&models.UserToken{
			UserId: foundUser.UID,
		}, models.UserToken{
			Token:        token,
			RefreshToken: refreshToken,
			UpdatedAt:    time.Now(),
		}).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while upserting tokens!"})
			return
		}

		loginResp := loginResponse{
			UserId: foundUser.UID,
			Token:  token,
		}

		c.JSON(http.StatusOK, loginResp)
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.CheckUserType(c, "ADMIN") // This can only be accessed by admins
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		recordPerPage, err := strconv.Atoi(c.DefaultQuery("recordPerPage", "10"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		page, err1 := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err1 != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage

		var allUsers []models.User
		err = db.Limit(recordPerPage).Offset(startIndex).Find(&allUsers).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while retrieving users!"})
			return
		}

		c.JSON(http.StatusOK, allUsers)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id") // c allows you to access parameters from Postman

		err := helper.MatchUserTypeToUid(c, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		err = db.Where("id = ?", userId).First(&user).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func generateUniqueUserId() string {
	// Generate a new UUID (version 4) and return it as a string
	// This will produce a unique identifier for each user.
	return uuid.New().String()
}
