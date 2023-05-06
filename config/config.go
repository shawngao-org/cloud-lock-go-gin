package config

import (
	"cloud-lock-go-gin/logger"
	"gopkg.in/yaml.v3"
	"os"
)

var Conf = getConfig()

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
	} `yaml:"security"`
}

func getConfig() Config {
	configFileName := "config.yml"
	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		logger.LogErr("[Config] Configuration file is not exist !")
		readFileErrLogImpl(configFileName, err)
		os.Exit(-1)
	}
	content, err := os.ReadFile(configFileName)
	if err != nil {
		readFileErrLogImpl(configFileName, err)
		os.Exit(-1)
	}

	var config Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		readFileErrLogImpl(configFileName, err)
		os.Exit(-1)
	}
	logger.LogSuccess("[Config] Configuration file '%s' -----> SUCCESS", configFileName)
	return config
}

func readFileErrLogImpl(fileName string, err error) {
	logger.LogErr("[Config] Configuration file '%s' -----> FAILED", fileName)
	logger.LogErr("%s", err)
	os.Exit(-1)
}
