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
	KaiHeiLa *KaiHeiLa `yaml:"KaiHeiLa"`
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
	BotAPILink         string `yaml:"BotAPILink"`
}
type KaiHeiLa struct {
	Switch bool   `yaml:"Switch"`
	Token  string `yaml:"Token"`
}

var (
	GetConfig  *Config
	ConfigPath string
)

//加载配置文件
func LoadConfig() error {
	//检查配置文件是否存在
	_, err := os.Lstat(ConfigPath)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		return err
	}
	newStu := &Config{}
	err = yaml.Unmarshal(content, &newStu)
	if err != nil {
		return err
	}
	GetConfig = newStu
	return nil
}
