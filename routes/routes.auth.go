package routes

import (
	"qna-management/handlers/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	AuthRepository := auth.InitAuthRepository(db)
	AuthService := auth.InitAuthService(AuthRepository)
	AuthHandler := auth.InitHandlerAuth(AuthService)

	// All auth route
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", AuthHandler.RegisterHandler)
	groupRoute.POST("/login", AuthHandler.LoginHandler)
}
