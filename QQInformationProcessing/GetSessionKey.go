package QQInformationProcessing

import (
	"encoding/json"
	"fmt"

	"xyz.nyan/MediaWiki-Bot/utils"
)

type verifyJson struct {
	Code    int    `json:"code"`
	Session string `json:"session"`
}

func GetSessionKey() string {
	Config := utils.ReadConfig()
	requestBody := fmt.Sprintf(`{
		"verifyKey": "%s"
	}`, Config.QQBot.VerifyKey)
	url := Config.QQBot.APILink + "/verify"
	body := utils.PostRequestJosn(url, requestBody)

	var config verifyJson
	json.Unmarshal([]byte(body), &config)
	SessionKey := config.Session

	//绑定Key与QQ
	requestBody = fmt.Sprintf(`{
		"sessionKey": "%s",
		"qq": %d
	}`, SessionKey, Config.QQBot.BotQQNumber)
	url = Config.QQBot.APILink + "/bind"
	utils.PostRequestJosn(url, requestBody)

	return SessionKey
}
