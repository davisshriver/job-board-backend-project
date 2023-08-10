package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/davisshriver/job-board-backend-project/database"
	helper "github.com/davisshriver/job-board-backend-project/helpers"
	"github.com/davisshriver/job-board-backend-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db = database.GetDB()
var validate = validator.New()

type loginResponse struct {
	UserId int    `json:"user_id"`
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

		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

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
		err = db.Where(models.UserToken{UserID: foundUser.UserID}).
			Assign(models.UserToken{
				Token:        token,
				RefreshToken: refreshToken,
				UpdatedAt:    time.Now(),
			}).
			FirstOrCreate(&models.UserToken{}).
			Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while upserting tokens!"})
			return
		}

		loginResp := loginResponse{
			UserId: foundUser.UserID,
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
		err = db.Where("user_id = ?", userId).First(&user).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var existingUser models.User

		err := helper.CheckUserType(c, "ADMIN") // This can only be accessed by admins
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userIDStr := c.Param("user_id")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = db.Where("user_id = ?", userID).First(&existingUser).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var updatedUser models.UserUpdate
		err = c.BindJSON(&updatedUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update only the non-null fields of the existing user with the updated values
		updateFields := map[string]interface{}{}

		if updatedUser.FirstName != nil {
			updateFields["first_name"] = *updatedUser.FirstName
		}
		if updatedUser.LastName != nil {
			updateFields["last_name"] = *updatedUser.LastName
		}
		if updatedUser.Password != nil {
			updateFields["password"] = *updatedUser.Password
		}
		if updatedUser.Email != nil {
			updateFields["email"] = *updatedUser.Email
		}
		if updatedUser.Phone != nil {
			updateFields["phone"] = *updatedUser.Phone
		}
		if updatedUser.UserType != nil {
			updateFields["user_type"] = *updatedUser.UserType
		}

		updateFields["updated_at"] = time.Now()

		// Don't perform update if there are no fields to update
		if len(updateFields) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No fields to update"})
			return
		}

		err = db.Model(&existingUser).Updates(updateFields).Error
		if err != nil {
			msg := fmt.Sprintf("User profile was not updated properly!")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, existingUser)
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.CheckUserType(c, "ADMIN")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userIdStr := c.Param("user_id")
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User

		// Attempt to retrieve the user with the given ID
		if err := db.Where("user_id = ?", userId).First(&user).Error; err != nil {
			if strings.Contains(err.Error(), "record not found") {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		// Delete the post
		if err := db.Delete(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "User deleted from the database"})
	}
}
