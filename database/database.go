package database

import (
	"cloud-lock-go-gin/config"
	"cloud-lock-go-gin/logger"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"sync"
)

var (
	DbMutex sync.Mutex
	Db      *gorm.DB
)

func ConnectDb() {
	host := config.Conf.Database.Mysql.Host
	port := config.Conf.Database.Mysql.Port
	user := config.Conf.Database.Mysql.User
	pwd := config.Conf.Database.Mysql.Password
	db := config.Conf.Database.Mysql.Db
	logger.LogInfo("Connecting to mysql server " + host + ":" + port + "...")
	uri := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + db
	dsn, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		logger.LogErr("Connection to mysql server " + host + ":" + port + "")
		logger.LogErr("%s", err)
		os.Exit(-1)
	}
	logger.LogSuccess("Connection to mysql server " + host + ":" + port + "")
	DbMutex.Lock()
	Db = dsn
	DbMutex.Unlock()
}
