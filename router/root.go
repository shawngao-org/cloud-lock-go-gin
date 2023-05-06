package router

import (
	"github.com/gin-gonic/gin"
)

func RootRouter(r *gin.Engine) {
	r.GET("/ping", ping)
	r.POST("/login", login)
}
