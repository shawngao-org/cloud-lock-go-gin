package middleware

import (
	"cloud-lock-go-gin/jwt"
	"cloud-lock-go-gin/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request pre logic code
		// next
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		jt, err := jwt.ParseToken(token)
		if err != nil {
			logger.LogErr("AuthMiddleware", "%s", err)
			c.AbortWithStatus(http.StatusForbidden)
			c.JSON(http.StatusForbidden, gin.H{
				"message":   "Invalid token format.",
				"exception": err.Error(),
			})
			return
		}
		err = jt.Claims.Valid()
		if err != nil {
			logger.LogErr("AuthMiddleware", "%s", err)
			c.AbortWithStatus(http.StatusForbidden)
			c.JSON(http.StatusForbidden, gin.H{
				"message":   "Invalid token.",
				"exception": err.Error(),
			})
			return
		}
		c.Next()
	}
}

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request after logic code
		// next
		c.Next()
	}
}
