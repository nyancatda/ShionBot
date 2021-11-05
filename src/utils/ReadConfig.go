package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Run  *Run        `yaml:"Run"`
	SNS  *SNS        `yaml:"SNS"`
	Wiki interface{} `yaml:"Wiki"`
}

type Run struct {
	WebHookPort string `yaml:"WebHookPort"`
	WebHookKey  string `yaml:"WebHookKey"`
	Language    string `yaml:"Language"`
}

type SNS struct {
	QQ       *QQ       `yaml:"QQ"`
	Telegram *Telegram `yaml:"Telegram"`
	Line     *Line     `yaml:"Line"`
}
type QQ struct {
	Switch      bool   `yaml:"Switch"`
	APILink     string `yaml:"APILink"`
	BotQQNumber int    `yaml:"BotQQNumber"`
	VerifyKey   string `yaml:"VerifyKey"`
}
type Telegram struct {
	Switch     bool   `yaml:"Switch"`
	Token      string `yaml:"Token"`
	BotAPILink string `yaml:"BotAPILink"`
}
type Line struct {
	Switch             bool   `yaml:"Switch"`
	ChannelAccessToken string `yaml:"ChannelAccessToken"`
}

func ReadConfig() *Config {
	content, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}
	newStu := &Config{}
	err = yaml.Unmarshal(content, &newStu)
	if err != nil {
		panic(err)
	}
	return newStu
}

func CheckConfigFile() bool {
	_, err := os.Lstat("./config.yml")
	return os.IsNotExist(err)
}
