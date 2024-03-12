package main

import (
	"tablelink/server"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	server.RegisterAPIService(r)
}

func main() {
	_ = r.Run()
}
