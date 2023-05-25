package controller

import (
	"cloud-lock-go-gin/config"
	"cloud-lock-go-gin/database"
	"cloud-lock-go-gin/jwt"
	"cloud-lock-go-gin/logger"
	"cloud-lock-go-gin/redis"
	"cloud-lock-go-gin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var pack = "UserController"

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong!",
	})
}

func Login(context *gin.Context) {
	username := context.PostForm("user")
	password := context.PostForm("password")
	_, err := database.GetUserByNameAndPwd(username, password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message":   "Internal server error",
			"exception": err.Error(),
		})
		return
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message":   "Internal server error",
			"exception": err,
		})
		return
	}
	err = redis.Client.Set(context,
		token, username, time.Duration(config.Conf.Security.Jwt.Timeout)*time.Second,
	).Err()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message":   "Internal server error",
			"exception": err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"token":   token,
	})
}

func Test(context *gin.Context) {
	plainText := "Hello world."
	logger.LogInfo(pack, "Plain text: %s", plainText)
	cipherText := util.Encrypted(config.Conf.Security.Rsa.Public, plainText)
	logger.LogInfo(pack, "Cipher text: %s", cipherText)
	decryptText := util.Decrypted(config.Conf.Security.Rsa.Private, cipherText)
	logger.LogInfo(pack, "Decrypt text: %s", decryptText)
	context.JSON(http.StatusOK, gin.H{})
}
