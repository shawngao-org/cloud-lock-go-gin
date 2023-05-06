package database

import (
	"cloud-lock-go-gin/config"
	"cloud-lock-go-gin/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var Conn = connectDb()

func connectDb() *sql.DB {
	host := config.Conf.Database.Mysql.Host
	port := config.Conf.Database.Mysql.Port
	user := config.Conf.Database.Mysql.User
	pwd := config.Conf.Database.Mysql.Password
	db := config.Conf.Database.Mysql.Db
	logger.LogInfo("[MYSQL] Connecting to mysql server " + host + ":" + port + "...")
	dsn, err := sql.Open("mysql", user+":"+pwd+"@tcp("+host+":"+port+")/"+db)
	if err != nil {
		logger.LogErr("[MYSQL] Connection to mysql server " + host + ":" + port + " -----> FAILED")
		logger.LogErr("%s", err)
		err := dsn.Close()
		if err != nil {
			logger.LogErr("[MYSQL] Close mysql connect " + host + ":" + port + " -----> FAILED")
			logger.LogErr("%s", err)
		}
		os.Exit(-1)
	}
	logger.LogSuccess("[MYSQL] Connection to mysql server " + host + ":" + port + " -----> SUCCESS")
	return dsn
}
