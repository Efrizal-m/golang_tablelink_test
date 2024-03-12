package server

import (
	"tablelink/internal/user"

	"github.com/gin-gonic/gin"
)

func registerUserRoute(r *gin.Engine, userController user.HTTPController) {
	userRouter := r.Group("/v1/users")
	userRouter.POST("/", userController.Add)
	userRouter.PUT("/:name", userController.Update)
	userRouter.DELETE("/:id", userController.Delete)
}
