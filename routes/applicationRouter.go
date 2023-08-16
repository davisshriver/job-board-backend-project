package routes

import (
	controller "github.com/davisshriver/job-board-backend-project/controllers"
	"github.com/davisshriver/job-board-backend-project/middleware"
	"github.com/gin-gonic/gin"
)

func ApplicationRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users/:user_id/applications", controller.GetUserApplications())
	incomingRoutes.GET("/users/:user_id/applications/:application_id", controller.GetApplication())
	incomingRoutes.POST("/users/:user_id/posts/:post_id/applications", controller.CreateApplication())
	//incomingRoutes.PATCH("/users/:user_id/applications/:application_id", controller.UpdateApplication())
	incomingRoutes.DELETE("/users/:user_id/applications/:application_id", controller.DeleteApplication())
}
