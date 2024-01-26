package routes

import (
	loginAuth "qna-management/controllers/auth-controllers/login"
	"qna-management/controllers/auth-controllers/register"
	handlerLogin "qna-management/handlers/auth-handlers/login"
	handlerRegister "qna-management/handlers/auth-handlers/register"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	// All handler auth
	LoginRepository := loginAuth.NewRepositoryLogin(db)
	LoginService := loginAuth.NewServiceLogin(LoginRepository)
	LoginHandler := handlerLogin.NewHandlerLogin(LoginService)

	RegisterRepository := register.NewRepositoryRegister(db)
	RegisterService := register.NewServiceRegister(RegisterRepository)
	RegisterHandler := handlerRegister.NewHandlerRegister(RegisterService)

	// All auth route
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/login", LoginHandler.LoginHandler)
	groupRoute.POST("/register", RegisterHandler.RegisterHandler)
}
