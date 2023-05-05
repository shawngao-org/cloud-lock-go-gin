package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var secrets = gin.H{
	"admin":  gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func AdminRouter(r *gin.Engine) {
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "123",
	}))
	authorized.GET("/secret", func(context *gin.Context) {
		user := context.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": "No secret :("})
		}
	})
}
