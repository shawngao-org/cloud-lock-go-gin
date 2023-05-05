package router

import "github.com/gin-gonic/gin"

func LoadRouter(r *gin.Engine) {
	RootRouter(r)
	AdminRouter(r)
}
