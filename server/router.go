package server

import (
	"github.com/gin-gonic/gin"
)

func registerUserRoute(r *gin.Engine, userController user.HTTPController) {
	userRouter := r.Group("/v1/users")
	userRouter.GET("/:id", userController.findAll)
	userRouter.POST("/", userController.Add)
	userRouter.PUT("/:name", userController.Update)
	userRouter.DELETE("/:id", userController.Delete)
}
