package main

import (
	"cloud-lock-go-gin/router"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var conf = getConfig()
var conn = connectDb(conf.Database.Mysql.Host, conf.Database.Mysql.Port,
	conf.Database.Mysql.User, conf.Database.Mysql.Password,
	conf.Database.Mysql.Db)

func main() {
	r := gin.Default()
	router.LoadRouter(r)
	srv := &http.Server{
		Addr:    conf.Server.Ip + ":" + conf.Server.Port,
		Handler: r,
	}
	go func() {
		startServer(srv)
	}()
	shutdownServer(srv)
}

func startServer(srv *http.Server) {
	logInfo("Start: Starting server...")
	logSuccess("Listen: Listening address -----> %s", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		logWarn("%s", err)
	}
}

func shutdownServer(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logWarn("Shutdown: Shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		logWarn("Shutdown: ")
		logWarn("%s", err)
	}
	logSuccess("Shutdown: Exited -----> SUCCESS")
}
