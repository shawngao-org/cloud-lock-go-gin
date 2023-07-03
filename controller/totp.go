package controller

import (
	"cloud-lock-go-gin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTotp godoc
// @Summary      Get TOTP secret
// @Tags         TOTP
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /totp/generateKey [get]
// @securityDefinitions.basic BasicAuth
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

// GenerateCode godoc
// @Summary      Get TOTP code
// @Tags         TOTP
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /totp/generateCode [get]
// @securityDefinitions.basic BasicAuth
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
