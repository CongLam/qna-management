package routes

import (
	"qna-management/handlers/auth"
	"qna-management/handlers/graph"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB, route *gin.Engine) {
	AuthRepository := auth.InitAuthRepository(db)
	AuthService := auth.InitAuthService(AuthRepository)
	AuthHandler := auth.InitHandlerAuth(AuthService)

	GraphRepository := graph.InitGraphRepository(db)
	GraphService := graph.InitGraphService(GraphRepository)
	GraphHandler := graph.InitHandlerGraph(GraphService)

	// All auth route
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", AuthHandler.RegisterHandler)
	groupRoute.POST("/login", AuthHandler.LoginHandler)
	// groupRoute.POST("/graph", AuthHandler.GetGraph)

	// Map route
	groupRoute.POST("/graph", GraphHandler.GetGraph)
}
