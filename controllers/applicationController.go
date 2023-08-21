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
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

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
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

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

		applicationModel := models.Application{
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
		}
		
		result := db.Create(&applicationModel)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create application"})
			return
		}
		
		output := outputs.ApplicationOutput{
			ApplicationID: applicationModel.ApplicationID,
			UserID:        applicationModel.UserID,
			PostID:        applicationModel.PostID,
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
			Referrals:     application.Referrals,
			DesiredSalary: application.DesiredSalary,
			Availability:  application.Availability,
			Education:     application.Education,
			WorkHistory:   application.WorkHistory,
			CreatedAt:     applicationModel.CreatedAt,
			UpdatedAt:     applicationModel.UpdatedAt,
		}

		c.JSON(http.StatusOK, output)
	}
}

func UpdateApplication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var existingApplication models.Application

		// Grab parameters
		userIdStr := c.Param("user_id")
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		applicationIdStr := c.Param("application_id")
		applicationId, err := strconv.Atoi(applicationIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the current user matches the user from parameters
		currentUserId := helpers.GetUserIdFromToken(c)
		if currentUserId != userId {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "You can only update your own applications"})
			return
		}

		// Retrieve application information
		err = db.Where("application_id = ? AND user_id = ?", applicationId, userId).First(&existingApplication).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var updatedInput inputs.ApplicationUpdateInput
		err = c.BindJSON(&updatedInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update fields in the existingApplication model based on updatedInput
		if updatedInput.FirstName != nil {
			existingApplication.FirstName = *updatedInput.FirstName
		}
		if updatedInput.LastName != nil {
			existingApplication.LastName = *updatedInput.LastName
		}
		if updatedInput.Email != nil {
			existingApplication.Email = *updatedInput.Email
		}
		if updatedInput.Phone != nil {
			existingApplication.Phone = *updatedInput.Phone
		}
		if updatedInput.Address != nil {
			existingApplication.Address = *updatedInput.Address
		}
		if updatedInput.City != nil {
			existingApplication.City = *updatedInput.City
		}
		if updatedInput.State != nil {
			existingApplication.State = *updatedInput.State
		}
		if updatedInput.PostalCode != nil {
			existingApplication.PostalCode = *updatedInput.PostalCode
		}
		if updatedInput.CoverLetter != nil {
			existingApplication.CoverLetter = *updatedInput.CoverLetter
		}
		if updatedInput.ResumeURL != nil {
			existingApplication.ResumeURL = *updatedInput.ResumeURL
		}
		if updatedInput.LinkedInURL != nil {
			existingApplication.LinkedInURL = *updatedInput.LinkedInURL
		}
		if updatedInput.PortfolioURL != nil {
			existingApplication.PortfolioURL = *updatedInput.PortfolioURL
		}
		if updatedInput.DesiredSalary != nil {
			existingApplication.DesiredSalary = *updatedInput.DesiredSalary
		}
		if updatedInput.Availability != nil {
			existingApplication.Availability = *updatedInput.Availability
		}

		existingApplication.UpdatedAt = time.Now()

		// Marshal the fields that were originally stored as JSON
		// Must be done this way, not compatible with reflection due to slice arrays
		if updatedInput.Education != nil {
			educationJSON, err := json.Marshal(updatedInput.Education)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal education array"})
				return
			}
			existingApplication.Education = educationJSON
		}

		if updatedInput.Referrals != nil {
			referralsJSON, err := json.Marshal(updatedInput.Referrals)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal referrals array"})
				return
			}
			existingApplication.Referrals = referralsJSON
		}

		if updatedInput.WorkHistory != nil {
			workHistoryJSON, err := json.Marshal(updatedInput.WorkHistory)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal work history array"})
				return
			}
			existingApplication.WorkHistory = workHistoryJSON
		}

		// Perform the update
		if err := db.Save(&existingApplication).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Create response struct
		updatedOutput := outputs.ApplicationOutput{
			ApplicationID: existingApplication.ApplicationID,
			UserID:        existingApplication.UserID,
			PostID:        existingApplication.PostID,
			FirstName:     existingApplication.FirstName,
			LastName:      existingApplication.LastName,
			Email:         existingApplication.Email,
			Phone:         existingApplication.Phone,
			Address:       existingApplication.Address,
			City:          existingApplication.City,
			State:         existingApplication.State,
			PostalCode:    existingApplication.PostalCode,
			CoverLetter:   existingApplication.CoverLetter,
			ResumeURL:     existingApplication.ResumeURL,
			LinkedInURL:   existingApplication.LinkedInURL,
			PortfolioURL:  existingApplication.PortfolioURL,
			DesiredSalary: existingApplication.DesiredSalary,
			Availability:  existingApplication.Availability,
			CreatedAt:     existingApplication.CreatedAt,
			UpdatedAt:     existingApplication.UpdatedAt,
		}

		// Unmarshal existing post's Experience, Education, and Work Hitory for output
		if err := json.Unmarshal(existingApplication.Referrals, &updatedOutput.Referrals); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal referrals array"})
			return
		}

		if err := json.Unmarshal(existingApplication.Education, &updatedOutput.Education); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal education array"})
			return
		}

		if err := json.Unmarshal(existingApplication.WorkHistory, &updatedOutput.WorkHistory); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal work history array"})
			return
		}

		c.JSON(http.StatusOK, updatedOutput)
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
