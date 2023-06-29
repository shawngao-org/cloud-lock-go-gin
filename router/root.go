package router

import (
	"cloud-lock-go-gin/controller"
	"github.com/gin-gonic/gin"
)

func RootRouter(r *gin.Engine) {
	r.POST("/login", controller.Login)
	r.GET("/ping", controller.Ping)
}
