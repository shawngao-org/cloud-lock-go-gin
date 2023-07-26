package controller

import (
	"cloud-lock-go-gin/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// GetPubKey godoc
// @Summary      Get RSA public key string
// @Tags         RSA
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /rsa/getPubKey [get]
func GetPubKey(context *gin.Context) {
	pk, err := os.ReadFile(config.Conf.Security.Rsa.Public)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message":   "failed to read public key file",
			"exception": err.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"message": string(pk),
	})
}
