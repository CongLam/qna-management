package routes

import (
	"qna-management/controllers/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	// All handler auth
	// LoginRepository := loginAuth.NewRepositoryLogin(db)
	// LoginService := loginAuth.NewServiceLogin(LoginRepository)
	// LoginHandler := handlerLogin.NewHandlerLogin(LoginService)

	RegisterRepository := auth.InitAuthRepository(db)
	RegisterService := auth.InitAuthService(RegisterRepository)
	RegisterHandler := auth.InitHandlerAuth(RegisterService)

	// All auth route
	groupRoute := route.Group("/api/v1")
	// groupRoute.POST("/login", LoginHandler.LoginHandler)
	groupRoute.POST("/register", RegisterHandler.RegisterHandler)
}
