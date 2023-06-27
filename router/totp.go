package router

import (
	"cloud-lock-go-gin/controller"
	"cloud-lock-go-gin/middleware"
	"github.com/gin-gonic/gin"
)

func TotpRouter(r *gin.Engine) {
	authorized := r.Group("/totp", middleware.AuthMiddleware())
	{
		authorized.POST("/generateKey", controller.GetTotp)
		authorized.POST("/generateCode", controller.GenerateCode)
	}
}
