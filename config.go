package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

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
	} `yaml:"database"`
}

func getConfig() Config {
	configFileName := "config.yml"
	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		logErr("Configuration file is not exist !")
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
	logSuccess("Read: configuration file '%s' -----> SUCCESS", configFileName)
	return config
}

func readFileErrLogImpl(fileName string, err error) {
	logErr("Read: configuration file '%s' -----> FAILED", fileName)
	logErr("%s", err)
	os.Exit(-1)
}
