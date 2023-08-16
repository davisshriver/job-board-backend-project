package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	inputs "github.com/davisshriver/job-board-backend-project/controllers/inputs"
	outputs "github.com/davisshriver/job-board-backend-project/controllers/outputs"
	helper "github.com/davisshriver/job-board-backend-project/helpers"
	helpers "github.com/davisshriver/job-board-backend-project/helpers"
	"github.com/davisshriver/job-board-backend-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		applicationId := c.Param("application_id")
		userId := c.Param("user_id")
		userIdInt, err := strconv.Atoi(userId)

		// Check if the current user matches the user from parameters
		currentUserId := helpers.GetUserIdFromToken(c)
		if currentUserId != userIdInt {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "You can only retrieve your own applications"})
			return
		}

		var application models.Application
		err = db.Where("application_id = ? AND user_id = ?", applicationId, userId).First(&application).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var education []inputs.EducationInfo
		var referrals []inputs.Referral
		var workHistory []inputs.WorkExperience

		err = json.Unmarshal(application.Education, &education)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal education array"})
			return
		}

		err = json.Unmarshal(application.Referrals, &referrals)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal referrals array"})
			return
		}

		err = json.Unmarshal(application.WorkHistory, &workHistory)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal work history array"})
			return
		}

		applicationResponse := outputs.ApplicationOutput{
			ApplicationID: application.ApplicationID,
			UserID:        application.UserID,
			PostID:        application.PostID,
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
			Referrals:     referrals,
			DesiredSalary: application.DesiredSalary,
			Availability:  application.Availability,
			Education:     education,
			WorkHistory:   workHistory,
			CreatedAt:     application.CreatedAt,
			UpdatedAt:     application.UpdatedAt,
		}

		c.JSON(http.StatusOK, applicationResponse)
	}
}

func GetUserApplications() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		userIdInt, err := strconv.Atoi(userId)

		// Check if the current user matches the user from parameters
		currentUserId := helpers.GetUserIdFromToken(c)
		if currentUserId != userIdInt {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "You can only retrieve your own applications"})
			return
		}

		var applications []models.Application

		err = db.Where("user_id = ?", userId).Find(&applications).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var applicationResponses []outputs.ApplicationOutput

		for _, application := range applications {
			var education []inputs.EducationInfo
			var referrals []inputs.Referral
			var workHistory []inputs.WorkExperience

			err = json.Unmarshal(application.Education, &education)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal education array"})
				return
			}

			err = json.Unmarshal(application.Referrals, &referrals)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal referrals array"})
				return
			}

			err = json.Unmarshal(application.WorkHistory, &workHistory)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal work history array"})
				return
			}

			applicationResponse := outputs.ApplicationOutput{
				ApplicationID: application.ApplicationID,
				UserID:        application.UserID,
				PostID:        application.PostID,
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
				Referrals:     referrals,
				DesiredSalary: application.DesiredSalary,
				Availability:  application.Availability,
				Education:     education,
				WorkHistory:   workHistory,
				CreatedAt:     application.CreatedAt,
				UpdatedAt:     application.UpdatedAt,
			}

			applicationResponses = append(applicationResponses, applicationResponse)
		}

		c.JSON(http.StatusOK, applicationResponses)
	}
}

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

func DeleteApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		applicationID := c.Param("application_id")

		err := helper.CheckUserType(c, "ADMIN")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var application models.Application
		err = db.Where("application_id = ?", applicationID).First(&application).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Delete the application
		if err := db.Delete(&application).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete application"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "Application deleted"})
	}
}
