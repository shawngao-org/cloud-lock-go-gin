package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func connectDb(host string, port string, user string, pwd string, db string) *sql.DB {
	logInfo("MYSQL: Connecting to mysql server " + host + ":" + port + "...")
	dsn, err := sql.Open("mysql", user+":"+pwd+"@tcp("+host+":"+port+")/"+db)
	if err != nil {
		logErr("MYSQL: Connection to mysql server " + host + ":" + port + " -----> FAILED")
		logErr("%s", err)
		err := dsn.Close()
		if err != nil {
			logErr("MYSQL: Close mysql connect " + host + ":" + port + " -----> FAILED")
			logErr("%s", err)
		}
		os.Exit(-1)
	}
	logSuccess("MYSQL: Connection to mysql server " + host + ":" + port + " -----> SUCCESS")
	return dsn
}
