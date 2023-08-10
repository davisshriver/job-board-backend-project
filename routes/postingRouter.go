package routes

import (
	controller "github.com/davisshriver/job-board-backend-project/controllers"
	"github.com/davisshriver/job-board-backend-project/middleware"
	"github.com/gin-gonic/gin"
)

func PostingRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate()) // Use middleware to secure routes via authentication
	incomingRoutes.GET("/posts", controller.GetJobPosts())
	incomingRoutes.GET("posts/:post_id", controller.GetJobPost())
	incomingRoutes.POST("/posts", controller.CreateJobPost())
	incomingRoutes.PATCH("/posts/:post_id", controller.UpdateJobPost())
	incomingRoutes.DELETE("posts/:post_id", controller.DeleteJobPost())
}
