package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	// All handler auth
	LoginRepository := loginAuth
}
