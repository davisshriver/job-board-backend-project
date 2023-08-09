package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
		postIdStr := c.Param("post_id") // c allows you to access parameters from Postman
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

func PostJob() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post models.JobPost

		err := helper.CheckUserType(c, "ADMIN") // This can only be accessed by admins
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

func DeleteJob() gin.HandlerFunc {
	return func(c *gin.Context) {

		err := helper.CheckUserType(c, "ADMIN") // This can only be accessed by admins
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		postIdStr := c.Param("post_id") // c allows you to access parameters from Postman
		postId, err := strconv.Atoi(postIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var post models.JobPost

		// Check if a post with that Id exists
		err = db.Where("post_id = ?", postId).First(&post).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = db.Delete(&post).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Success": "Entry deleted from database!"})
	}
}
