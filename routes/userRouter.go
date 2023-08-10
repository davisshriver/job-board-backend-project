package routes

import (
	controller "github.com/davisshriver/job-board-backend-project/controllers"
	"github.com/davisshriver/job-board-backend-project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate()) // Use middleware to secure routes via authentication
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.PATCH("/users/:user_id", controller.UpdateUser())
	incomingRoutes.DELETE("/users/:user_id", controller.DeleteUser())
}
