package router

import (
	"cloud-lock-go-gin/config"
	"cloud-lock-go-gin/controller"
	"cloud-lock-go-gin/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files

func RootRouter(r *gin.Engine) {
	r.POST("/login", controller.Login)
	r.GET("/ping", controller.Ping)
	docs.SwaggerInfo.Title = "API Docs"
	docs.SwaggerInfo.Description = "null."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.Conf.Server.Ip + ":" + config.Conf.Server.Port
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
