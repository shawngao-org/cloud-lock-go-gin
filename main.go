package main

import (
	"cloud-lock-go-gin/config"
	"cloud-lock-go-gin/logger"
	"cloud-lock-go-gin/router"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(requestMiddleware())
	r.Use(responseMiddleware())
	router.LoadRouter(r)
	srv := &http.Server{
		Addr:    config.Conf.Server.Ip + ":" + config.Conf.Server.Port,
		Handler: r,
	}
	go func() {
		startServer(srv)
	}()
	shutdownServer(srv)
}

func startServer(srv *http.Server) {
	logger.LogInfo("[Server] Starting server...")
	logger.LogSuccess("[Server] Listening address -----> %s", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		logger.LogWarn("%s", err)
	}
}

func requestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request pre logic code

		// next
		c.Next()
	}
}

func responseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request after logic code

		// next
		c.Next()
	}
}

func shutdownServer(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.LogWarn("[Server] Shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		logger.LogWarn("[Server] ")
		logger.LogWarn("%s", err)
	}
	logger.LogSuccess("[Server] Exited -----> SUCCESS")
}
