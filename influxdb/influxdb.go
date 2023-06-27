package influxdb

import (
	"cloud-lock-go-gin/config"
	"cloud-lock-go-gin/logger"
	"context"
	_ "github.com/go-sql-driver/mysql"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"strconv"
	"sync"
	"time"
)

var (
	InfluxMutex sync.Mutex
	Influx      influxdb2.Client
)

func ConnectInfluxDb() {
	if !config.Conf.Database.Influxdb.Enable {
		return
	}
	host := config.Conf.Database.Influxdb.Host
	port := config.Conf.Database.Influxdb.Port
	token := config.Conf.Database.Influxdb.Token
	logger.LogInfo("Connecting to InfluxDb server " + host + ":" + port + "...")
	uri := "http://" + host + ":" + port
	client := influxdb2.NewClient(uri, token)
	logger.LogSuccess("Connection to InfluxDb server " + host + ":" + port)
	InfluxMutex.Lock()
	Influx = client
	InfluxMutex.Unlock()
}

func WriteReqLog(uid int64, path string, method string, ip string) {
	if !config.Conf.Database.Influxdb.Enable {
		return
	}
	org := config.Conf.Database.Influxdb.Org
	bucket := config.Conf.Database.Influxdb.Bucket
	writeAPI := Influx.WriteAPIBlocking(org, bucket)
	tags := map[string]string{
		"uid": strconv.FormatInt(uid, 10),
	}
	fields := map[string]interface{}{
		"event": "[" + method + "] [" + ip + "] " + path,
	}
	point := write.NewPoint("req_log", tags, fields, time.Now())
	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		logger.LogErr("%s", err.Error())
	}
}
