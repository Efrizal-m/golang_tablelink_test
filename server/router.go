package server

import (
	"github.com/gin-gonic/gin"
)

func registerUserRoute(r *gin.Engine, movieController user.HTTPController) {
	movieRouter := r.Group("/v1/users")
	movieRouter.GET("/", movieController.FindByTitle)
	movieRouter.GET("/:id", movieController.FindByID)
	movieRouter.POST("/", movieController.Add)
	movieRouter.PUT("/:id", movieController.Update)
	movieRouter.DELETE("/:id", movieController.Delete)
}
