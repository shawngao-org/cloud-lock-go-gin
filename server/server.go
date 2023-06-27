package server

import (
	"cloud-lock-go-gin/config"
	"cloud-lock-go-gin/logger"
	"cloud-lock-go-gin/middleware"
	"cloud-lock-go-gin/router"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func HandleServer() {
	r := gin.Default()
	r.Use(middleware.RequestMiddleware())
	r.Use(middleware.ResponseMiddleware())
	router.LoadRouter(r)
	for _, info := range r.Routes() {
		logger.LogInfo("[" + info.Method + "] => [" + info.Path + "]")
	}
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
	logger.LogInfo("Starting server...")
	logger.LogStarted()
	logger.LogSuccess("Listening address: %s", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		logger.LogWarn("%s", err)
	}
}

func shutdownServer(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.LogWarn("Shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		logger.LogWarn("%s", err)
	}
	logger.LogSuccess("Exited")
}
