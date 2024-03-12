package server

import (
	"tablelink/config"

	"tablelink/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func RegisterAPIService(e *gin.Engine) {
	db = config.GetDBConnection()

	registerUserAPIService(e)
}

func registerUserAPIService(r *gin.Engine) {
	userRepo := user.NewRepository(db)
	userUseCase := user.NewUseCase(userRepo)
	userController := user.NewHTTPController(userUseCase)

	// Start API
	registerUserRoute(r, userController)
}
