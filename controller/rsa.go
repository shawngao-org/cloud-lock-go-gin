package controller

import (
	"cloud-lock-go-gin/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPubKey(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": config.Conf.Security.Rsa.Public,
	})
}
