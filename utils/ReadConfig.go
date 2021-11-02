package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Run *Run `yaml:"Run"`
	QQBot *QQBot `yaml:"QQBot"`
	Wiki interface{} `yaml:"Wiki"`
}

type Run struct {
	WebHookPort string `yaml:"WebHookPort"`
	Language string `yaml:"Language"`
}

type QQBot struct {
	APILink string `yaml:"APILink"`
	BotQQNumber int `yaml:"BotQQNumber"`
	VerifyKey string `yaml:"VerifyKey"`
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