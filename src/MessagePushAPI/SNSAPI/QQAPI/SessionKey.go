package QQAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"xyz.nyan/ShionBot/src/utils"
)

type verifyJson struct {
	Code    int    `json:"code"`
	Session string `json:"session"`
}

func CreateSessionKey() (string, *http.Response, error) {
	//释放旧的SessionKey
	bytes, _ := ioutil.ReadFile("data/SessionKey")
	OldSessionKey := string(bytes)
	Config := utils.GetConfig
	requestBody := fmt.Sprintf(`{
		"verifyKey": "%s",
		"qq": %d
	  }`, OldSessionKey, Config.SNS.QQ.BotQQNumber)
	url := Config.SNS.QQ.APILink + "/release"
	utils.PostRequestJosn(url, requestBody)

	var SessionKey string
	Config = utils.GetConfig
	requestBody = fmt.Sprintf(`{
		"verifyKey": "%s"
	}`, Config.SNS.QQ.VerifyKey)
	url = Config.SNS.QQ.APILink + "/verify"
	body, resp, http_error := utils.PostRequestJosn(url, requestBody)

	var config verifyJson
	json.Unmarshal([]byte(body), &config)
	SessionKey = config.Session

	//绑定Key与QQ
	requestBody = fmt.Sprintf(`{
		"sessionKey": "%s",
		"qq": %d
	}`, SessionKey, Config.SNS.QQ.BotQQNumber)
	url = Config.SNS.QQ.APILink + "/bind"
	utils.PostRequestJosn(url, requestBody)

	//缓存SessionKey
	ioutil.WriteFile("data/SessionKey", []byte(SessionKey), 0664)
	return SessionKey, resp, http_error
}

func GetSessionKey() string {
	bytes, _ := ioutil.ReadFile("data/SessionKey")
	SessionKey := string(bytes)

	return SessionKey
}
