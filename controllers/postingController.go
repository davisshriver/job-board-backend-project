package controllers

import (
	"net/http"
	"strconv"

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
