package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RootRouter(r *gin.Engine) {
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})
}
