package middleware

import (
	"cloud-lock-go-gin/database"
	"cloud-lock-go-gin/logger"
	"cloud-lock-go-gin/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func RequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request pre logic code
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Conf.Develop {
			c.Next()
			return
		}
		token := c.Request.Header.Get("token")
		jt, err := util.ParseToken(token)
		if err != nil {
			logger.LogErr("%s", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":   "Invalid token format.",
				"exception": err.Error(),
			})
			return
		}
		err = jt.Claims.Valid()
		if err != nil {
			logger.LogErr("%s", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":   "Invalid token.",
				"exception": err.Error(),
			})
			return
		}
		uid := int64(jt.Claims.(jwt.MapClaims)["user"].(float64))
		r, e := database.CheckRouterPermission(c.Request.URL.Path, uid)
		if e != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{
				"exception": e.Error(),
			})
			return
		}
		if !r {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{})
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
