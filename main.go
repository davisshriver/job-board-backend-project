package main

import (
	"log"
	"os"

	"github.com/davisshriver/job-board-backend-project/database"
	"github.com/davisshriver/job-board-backend-project/models"
	"github.com/davisshriver/job-board-backend-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// Initialize the GORM database connection
	db := database.GetDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database connection")
	}
	defer sqlDB.Close()

	// Auto migrate the User model
	err = db.AutoMigrate(&models.User{}, &models.UserToken{}, &models.JobPost{})
	if err != nil {
		log.Fatal("failed to automigrate models")
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.PostingRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
