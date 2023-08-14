package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	inputs "github.com/davisshriver/job-board-backend-project/controllers/inputs"
	"github.com/davisshriver/job-board-backend-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var application inputs.ApplicationInput

		userIdStr := c.Param("user_id")
		userId, uIdErr := strconv.Atoi(userIdStr)

		postIdStr := c.Param("post_id")
		postId, postIdErr := strconv.Atoi(postIdStr)

		if uIdErr != nil || postIdErr != nil {
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

		educationJSON, err := json.Marshal(application.Education)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal education array"})
			return
		}

		referralsJSON, err := json.Marshal(application.Referrals)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal referrals array"})
			return
		}

		workHistoryJSON, err := json.Marshal(application.WorkHistory)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal work history array"})
			return
		}

		if err := db.Create(&models.Application{
			UserID:        userId,
			PostID:        postId,
			FirstName:     application.FirstName,
			LastName:      application.LastName,
			Email:         application.Email,
			Phone:         application.Phone,
			Address:       application.Address,
			City:          application.City,
			State:         application.State,
			PostalCode:    application.PostalCode,
			CoverLetter:   application.CoverLetter,
			ResumeURL:     application.ResumeURL,
			LinkedInURL:   application.LinkedInURL,
			PortfolioURL:  application.PortfolioURL,
			DesiredSalary: application.DesiredSalary,
			Availability:  application.Availability,
			Education:     educationJSON,
			Referrals:     referralsJSON,
			WorkHistory:   workHistoryJSON,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create application"})
			return
		}
		c.JSON(http.StatusOK, application)
	}
}
