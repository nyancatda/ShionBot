/*
 * @Author: NyanCatda
 * @Date: 2021-10-03 04:14:10
 * @LastEditTime: 2021-12-31 11:24:45
 * @LastEditors: NyanCatda
 * @Description: 读取配置文件
 * @FilePath: \ShionBot\src\utils\ReadConfig.go
 */
package utils

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Run struct {
		WebHookPort string `yaml:"WebHookPort"`
		WebHookKey  string `yaml:"WebHookKey"`
		Language    string `yaml:"Language"`
	} `yaml:"Run"`
	SNS struct {
		QQ struct {
			Switch      bool   `yaml:"Switch"`
			APILink     string `yaml:"APILink"`
			BotQQNumber int    `yaml:"BotQQNumber"`
			VerifyKey   string `yaml:"VerifyKey"`
		} `yaml:"QQ"`
		Telegram struct {
			Switch     bool   `yaml:"Switch"`
			Token      string `yaml:"Token"`
			BotAPILink string `yaml:"BotAPILink"`
		} `yaml:"Telegram"`
		Line struct {
			Switch             bool   `yaml:"Switch"`
			ChannelAccessToken string `yaml:"ChannelAccessToken"`
			BotAPILink         string `yaml:"BotAPILink"`
		} `yaml:"Line"`
		KaiHeiLa struct {
			Switch bool   `yaml:"Switch"`
			Token  string `yaml:"Token"`
		} `yaml:"KaiHeiLa"`
	} `yaml:"SNS"`
	Wiki interface{} `yaml:"Wiki"`
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
