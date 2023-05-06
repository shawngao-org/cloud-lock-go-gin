package router

import (
	"cloud-lock-go-gin/config"
	"cloud-lock-go-gin/jwt"
	"cloud-lock-go-gin/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong!",
	})
}

func login(context *gin.Context) {
	username := context.PostForm("user")
	//password := context.PostForm("password")
	token, err := jwt.GenerateToken(username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message":   "Internal server error",
			"exception": err,
		})
		return
	}
	err = redis.RedisClient.Set(context,
		username, token, time.Duration(config.Conf.Security.Jwt.Timeout),
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
