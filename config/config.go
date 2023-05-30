package config

import (
	"cloud-lock-go-gin/logger"
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"syscall"
	"time"
)

var (
	configMutex sync.Mutex
	Conf        Config
)

var pack = "config"

type Config struct {
	Server struct {
		Ip   string `yaml:"ip"`
		Port string `yaml:"port"`
	}
	Database struct {
		Mysql struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Db       string `yaml:"db"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
		} `yaml:"mysql"`
		Redis struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Db       int    `yaml:"db"`
			Password string `yaml:"password"`
		} `yaml:"redis"`
	} `yaml:"database"`
	Security struct {
		Jwt struct {
			Secret  string `yaml:"secret"`
			Timeout int64  `yaml:"timeout"`
		} `yaml:"jwt"`
		Rsa struct {
			Public  string `yaml:"public"`
			Private string `yaml:"private"`
		} `yaml:"rsa"`
	} `yaml:"security"`
	Nacos struct {
		Enable    bool   `yaml:"enable"`
		Ip        string `yaml:"ip"`
		Port      uint64 `yaml:"port"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		Namespace string `yaml:"namespace"`
		Group     string `yaml:"group"`
		DataId    string `yaml:"dataId"`
		Timeout   uint64 `yaml:"timeout"`
		Loglevel  string `yaml:"loglevel"`
	} `yaml:"nacos"`
	Develop bool `yaml:"develop"`
}

func parseContent2Config(content []byte) Config {
	var config Config
	err := yaml.Unmarshal(content, &config)
	if err != nil {
		readFileErrLogImpl(err)
		os.Exit(-1)
	}
	return config
}

func GetConfig() {
	configFileName := "config.yml"
	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		logger.LogErr(pack, "Configuration file is not exist !")
		readFileErrLogImpl(err)
		os.Exit(-1)
	}
	content, err := os.ReadFile(configFileName)
	if err != nil {
		readFileErrLogImpl(err)
		os.Exit(-1)
	}
	config := parseContent2Config(content)
	logger.LogSuccess(pack, "Configuration file '%s' -----> SUCCESS", configFileName)
	if config.Develop {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	if config.Nacos.Enable {
		logger.LogInfo(pack, "Configuration file mode: Nacos unified configuration center")
		nacosMain(config)
		return
	}
	logger.LogInfo(pack, "Profile Mode: Local config file")
	configMutex.Lock()
	Conf = config
	configMutex.Unlock()
}

func readFileErrLogImpl(err error) {
	logger.LogErr(pack, "Configuration file '%s' -----> FAILED")
	logger.LogErr(pack, "%s", err)
	os.Exit(-1)
}

func nacosMain(config Config) {
	sc := []constant.ServerConfig{{
		IpAddr: config.Nacos.Ip,
		Port:   config.Nacos.Port,
	}}
	cc := constant.ClientConfig{
		NamespaceId:         config.Nacos.Namespace,
		TimeoutMs:           config.Nacos.Timeout,
		NotLoadCacheAtStart: true,
		LogDir:              "log",
		LogLevel:            config.Nacos.Loglevel,
		Username:            config.Nacos.Username,
		Password:            config.Nacos.Password,
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		logger.LogErr(pack, "%s", err)
		os.Exit(-1)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: config.Nacos.DataId,
		Group:  config.Nacos.Group,
	})
	if err != nil {
		logger.LogErr(pack, "%s", err)
		os.Exit(-1)
	}
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: config.Nacos.DataId,
		Group:  config.Nacos.Group,
		OnChange: func(namespace, group, dataId, data string) {
			logger.LogInfo(pack, "The configuration file has changed...")
			logger.LogInfo(pack, "Group: %s, Data Id: %s", group, dataId)
			configMutex.Lock()
			Conf = parseContent2Config([]byte(data))
			configMutex.Unlock()
			if runtime.GOOS == "linux" {
				logger.LogWarn(pack, "The server is restarting, please wait a few seconds...")
				cmd := exec.Command("sh", "-c", "sleep 3 && exit 0")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Start(); err != nil {
					logger.LogErr(pack, "Error starting command: %s", err)
					os.Exit(1)
				}
				if err := cmd.Wait(); err != nil {
					logger.LogErr(pack, "Command finished with error: %s", err)
				}
				time.Sleep(1 * time.Second)
				binary, err := exec.LookPath(os.Args[0])
				if err != nil {
					logger.LogErr(pack, "Can't get executable binary.")
					logger.LogErr(pack, "Need to manually restart the server.")
					logger.LogErr(pack, "%s", err)
				}
				err = syscall.Exec(binary, os.Args, os.Environ())
				if err != nil {
					logger.LogErr(pack, "Need to manually restart the server.")
					logger.LogErr(pack, "%s", err)
				}
			} else {
				logger.LogErr(pack, "Need to manually restart the server.")
				logger.LogErr(pack, "Config hot reload is not support Windows OS.")
			}
		},
	})
	configMutex.Lock()
	Conf = parseContent2Config([]byte(content))
	configMutex.Unlock()
}
