package main

import (
	"qna-management/configs"
	"qna-management/routes"
	"qna-management/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Get the logger instance
	log := utils.GetLogger()

	// --------- Start application ----------
	log.Info("Starting the application...")

	// Setup router
	router := SetupRouter()

	// Run server
	log.Fatal(router.Run(":" + utils.GodotEnv("GO_PORT")))

	// --------- End program ----------
	log.Info("Ending the application...")
}

func SetupRouter() *gin.Engine {
	// Setup database connection
	db := configs.Connection()

	// Init router
	router := gin.Default()

	// Setup mode application
	if utils.GodotEnv("GO_ENV") != "production" && utils.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if utils.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Setup Middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	// Init all routes
	routes.InitAuthRoutes(db, router)

	return router
}
