package middleware

import (
	"github.com/gin-gonic/gin"
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
		//user, err := getCurrentUser(c)
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
