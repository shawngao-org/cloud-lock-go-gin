package router

import (
	"cloud-lock-go-gin/controller"
	"cloud-lock-go-gin/middleware"
	"github.com/gin-gonic/gin"
)

func TotpRouter(r *gin.Engine) {
	authorized := r.Group("/totp", middleware.AuthMiddleware())
	{
		authorized.GET("/generateKey", controller.GetTotp)
		authorized.GET("/generateCode", controller.GenerateCode)
	}
}
