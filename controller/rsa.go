package controller

import (
	"cloud-lock-go-gin/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetPubKey godoc
// @Summary      Get RSA public key string
// @Tags         RSA
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /rsa/getPubKey [get]
func GetPubKey(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": config.Conf.Security.Rsa.Public,
	})
}
