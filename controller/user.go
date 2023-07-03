package controller

import (
	"cloud-lock-go-gin/config"
	"cloud-lock-go-gin/database"
	"cloud-lock-go-gin/redis"
	"cloud-lock-go-gin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Ping godoc
// @Summary      Ping pong
// @Description  Ping-Pong
// @Tags         Ping
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /ping [get]
func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong!",
	})
}

// Login godoc
// @Summary      Login
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Param        email query string true "Email"
// @Param        password query string true "Password(RSA Encrypted)"
// @Router       /login [post]
func Login(context *gin.Context) {
	email := context.PostForm("email")
	password := context.PostForm("password")
	password = util.Decrypted(config.Conf.Security.Rsa.Private, password)
	user, err := database.GetUserByEmailAndPwd(email, password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message":   "Internal server error",
			"exception": err.Error(),
		})
		return
	}
	token, err := util.GenerateToken(user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message":   "Internal server error",
			"exception": err,
		})
		return
	}
	err = redis.Client.Set(context,
		strconv.FormatInt(user.Id, 10),
		token,
		time.Duration(config.Conf.Security.Jwt.Timeout)*time.Second,
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
