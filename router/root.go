package router

import (
	"cloud-lock-go-gin/controller"
	"cloud-lock-go-gin/middleware"
	"github.com/gin-gonic/gin"
)

func RootRouter(r *gin.Engine) {
	authorized := r.Group("/", middleware.AuthMiddleware())
	{
		authorized.GET("/ping", controller.Ping)
	}
	r.POST("/login", controller.Login)
	r.POST("/test", controller.Test)
}
