package main

import (
	"cloud-lock-go-gin/config"
	"cloud-lock-go-gin/logger"
	"cloud-lock-go-gin/middleware"
	"cloud-lock-go-gin/router"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var pack = "main"

func main() {
	r := gin.Default()
	r.Use(middleware.RequestMiddleware())
	r.Use(middleware.ResponseMiddleware())
	//r.Use(middleware.AuthMiddleware())
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
	logger.LogInfo(pack, "Starting server...")
	logger.LogSuccess(pack, "Listening address -----> %s", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		logger.LogWarn(pack, "%s", err)
	}
}

func shutdownServer(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.LogWarn(pack, "Shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		logger.LogWarn(pack, "%s", err)
	}
	logger.LogSuccess(pack, "Exited -----> SUCCESS")
}
