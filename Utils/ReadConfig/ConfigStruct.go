/*
 * @Author: NyanCatda
 * @Date: 2022-01-24 19:35:15
 * @LastEditTime: 2022-01-24 19:35:16
 * @LastEditors: NyanCatda
 * @Description: 配置文件结构体
 * @FilePath: \ShionBot\src\Utils\ReadConfig\ConfigStruct.go
 */
package ReadConfig

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