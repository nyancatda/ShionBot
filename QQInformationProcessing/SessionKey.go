package QQInformationProcessing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"xyz.nyan/MediaWiki-Bot/utils"
)

type verifyJson struct {
	Code    int    `json:"code"`
	Session string `json:"session"`
}

func CreateSessionKey() (string,*http.Response,error) {
	var SessionKey string
	Config := utils.ReadConfig()
	requestBody := fmt.Sprintf(`{
		"verifyKey": "%s"
	}`, Config.QQBot.VerifyKey)
	url := Config.QQBot.APILink + "/verify"
	body,resp,http_error := utils.PostRequestJosn(url, requestBody)

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
	return SessionKey,resp,http_error
}

func GetSessionKey() string {
    bytes, _ := ioutil.ReadFile("SessionKey")
    SessionKey := string(bytes)

	return SessionKey
}
