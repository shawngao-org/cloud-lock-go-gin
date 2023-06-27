package controller

import (
	"cloud-lock-go-gin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTotp(context *gin.Context) {
	secret, err := util.Totp()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"secret":  secret,
	})
}

func GenerateCode(context *gin.Context) {
	secret := context.PostForm("secret")
	code, err := util.GenerateCode(secret)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"codes":   code,
	})
}
