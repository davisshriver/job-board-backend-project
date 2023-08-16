package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	inputs "github.com/davisshriver/job-board-backend-project/controllers/inputs"
	helper "github.com/davisshriver/job-board-backend-project/helpers"
	"github.com/davisshriver/job-board-backend-project/models"
	"github.com/gin-gonic/gin"
)

func GetJobPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		recordPerPage, err := strconv.Atoi(c.DefaultQuery("recordPerPage", "10"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		page, err1 := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err1 != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage

		var allJobs []models.JobPost
		err = db.Limit(recordPerPage).Offset(startIndex).Find(&allJobs).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while retrieving job posts!"})
			return
		}

		c.JSON(http.StatusOK, allJobs)
	}
}

func GetJobPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		postIdStr := c.Param("post_id")
		postId, err := strconv.Atoi(postIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var post models.JobPost
		err = db.Where("post_id = ?", postId).First(&post).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, post)
	}
}

func CreateJobPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post models.JobPost

		err := helper.CheckUserType(c, "ADMIN")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = c.BindJSON(&post)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(post)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		post.CreatedAt = time.Now()

		// Create the user in the database using GORM
		err = db.Create(&post).Error
		if err != nil {
			msg := fmt.Sprintf("Job post was not created properly!")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, post)
	}
}

func UpdateJobPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var existingPost models.JobPost

		err := helper.CheckUserType(c, "ADMIN")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		postIdStr := c.Param("post_id")
		postId, err := strconv.Atoi(postIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = db.Where("post_id = ?", postId).First(&existingPost).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var updatedPost inputs.JobPostUpdate
		err = c.BindJSON(&updatedPost)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update only the non-null fields of the existing post with the updated values
		updateFields := map[string]interface{}{}

		if updatedPost.Role != nil {
			updateFields["role"] = *updatedPost.Role
		}
		if updatedPost.Description != nil {
			updateFields["description"] = *updatedPost.Description
		}
		if updatedPost.Requirements != nil {
			updateFields["requirements"] = *updatedPost.Requirements
		}
		if updatedPost.Wage != nil {
			updateFields["wage"] = *updatedPost.Wage
		}
		if updatedPost.Expires_At != nil {
			updateFields["expires_at"] = *updatedPost.Expires_At
		}

		// Don't perform update if there are no fields to update
		if len(updateFields) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No fields to update"})
			return
		}

		err = db.Model(&existingPost).Updates(updateFields).Error
		if err != nil {
			msg := fmt.Sprintf("Job post was not updated properly!")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, existingPost)
	}
}

func DeleteJobPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.CheckUserType(c, "ADMIN")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		postIdStr := c.Param("post_id")
		postId, err := strconv.Atoi(postIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var post models.JobPost

		// Attempt to retrieve the post with the given ID
		if err := db.Where("post_id = ?", postId).First(&post).Error; err != nil {
			if strings.Contains(err.Error(), "record not found") {
				c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		if err := db.Delete(&post).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "Post deleted from the database"})
	}
}
