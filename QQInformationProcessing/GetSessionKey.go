package QQInformationProcessing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"

	"xyz.nyan/MediaWiki-Bot/utils"
)

type verifyJson struct {
	Code    int    `json:"code"`
	Session string `json:"session"`
}

func GetSessionKey() string {
    bytes, _ := ioutil.ReadFile("SessionKey")
    RKey := string(bytes)

	fetchMessageUrl := utils.ReadConfig().QQBot.APILink + "/fetchMessage?sessionKey=" + RKey + "&count=1"
	fetchMessageBody := utils.HttpRequest(fetchMessageUrl)
	info := make(map[string]interface{})
	json.Unmarshal([]byte(fetchMessageBody), &info)

	var SessionKey string
	if int(math.Floor(info["code"].(float64))) == 3 {
		Config := utils.ReadConfig()
		requestBody := fmt.Sprintf(`{
			"verifyKey": "%s"
		}`, Config.QQBot.VerifyKey)
		url := Config.QQBot.APILink + "/verify"
		body := utils.PostRequestJosn(url, requestBody)

		var config verifyJson
		json.Unmarshal([]byte(body), &config)
		SessionKey = config.Session

		//绑定Key与QQ
		requestBody = fmt.Sprintf(`{
			"sessionKey": "%s",
			"qq": %d
		}`, SessionKey, Config.QQBot.BotQQNumber)
		url = Config.QQBot.APILink + "/bind"
		utils.PostRequestJosn(url, requestBody)

		//缓存SessionKey
		ioutil.WriteFile("SessionKey", []byte(SessionKey), 0664)
	} else {
		SessionKey = RKey
	}

	return SessionKey
}
