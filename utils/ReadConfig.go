package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	QQBot *QQBot `yaml:"QQBot"`
	Wiki *Wiki `yaml:"Wiki"`
}

type QQBot struct {
	APILink string `yaml:"APILink"`
	BotQQNumber int `yaml:"BotQQNumber"`
	VerifyKey string `yaml:"VerifyKey"`
}

type Wiki struct {
	WikiLink string `yaml:"WikiLink"`
}

func ReadConfig() (*Config) {
	content,err := ioutil.ReadFile("./config.yml")
	if err != nil{
		panic(err)
	}
	newStu := &Config{}
	err = yaml.Unmarshal(content,&newStu)
	if err != nil{
		panic(err)
	}
	return newStu
}

func CheckConfigFile() bool {
	_, err := os.Lstat("./config.yml")
	return os.IsNotExist(err)
}