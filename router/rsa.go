package router

import (
	"cloud-lock-go-gin/controller"
	"github.com/gin-gonic/gin"
)

func GetRsaPubKey(r *gin.Engine) {
	r.GET("/rsa/getPubKey", controller.GetPubKey)
}
