package controllers

import (
	"net/http"
	"strconv"
	"time"

	//helper "github.com/davisshriver/job-board-backend-project/helpers"
	"github.com/davisshriver/job-board-backend-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var application models.Application

		userIdStr := c.Param("user_id")
		userId, uIdErr := strconv.Atoi(userIdStr)
		
		postIdStr := c.Param("post_id")
		postId, postIdErr := strconv.Atoi(postIdStr)

		if uIdErr != nil || postIdErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error converting userId or postId to integer value!"})
			return
		}

		// Bind and validate JSON data
		if err := c.ShouldBindJSON(&application); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the struct using validator
		validate := validator.New()
		if err := validate.Struct(application); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		application.CreatedAt = time.Now()

		// Associate the application with the user (set the UserID field)
		application.UserID = userId
		application.PostID = postId

		// Create the application in the database using GORM
		if err := db.Create(&application).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create application"})
			return
		}

		c.JSON(http.StatusOK, application)
	}
}
