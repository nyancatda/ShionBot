/*
 * @Author: NyanCatda
 * @Date: 2021-11-05 13:51:15
 * @LastEditTime: 2022-01-24 19:50:01
 * @LastEditors: NyanCatda
 * @Description: Session处理API
 * @FilePath: \ShionBot\src\MessagePushAPI\SNSAPI\QQAPI\SessionKey.go
 */
package QQAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nyancatda/ShionBot/src/Utils/HttpRequest"
	"github.com/nyancatda/ShionBot/src/Utils/ReadConfig"
)

type verifyJson struct {
	Code    int    `json:"code"`
	Session string `json:"session"`
}

func CreateSessionKey() (string, *http.Response, error) {
	//释放旧的SessionKey
	bytes, _ := ioutil.ReadFile("data/SessionKey")
	OldSessionKey := string(bytes)
	Config := ReadConfig.GetConfig
	requestBody := fmt.Sprintf(`{
		"verifyKey": "%s",
		"qq": %d
	  }`, OldSessionKey, Config.SNS.QQ.BotQQNumber)
	url := Config.SNS.QQ.APILink + "/release"
	HttpRequest.PostRequestJson(url, requestBody, []string{})

	var SessionKey string
	Config = ReadConfig.GetConfig
	requestBody = fmt.Sprintf(`{
		"verifyKey": "%s"
	}`, Config.SNS.QQ.VerifyKey)
	url = Config.SNS.QQ.APILink + "/verify"
	body, resp, http_error := HttpRequest.PostRequestJson(url, requestBody, []string{})

	var config verifyJson
	json.Unmarshal([]byte(body), &config)
	SessionKey = config.Session

	//绑定Key与QQ
	requestBody = fmt.Sprintf(`{
		"sessionKey": "%s",
		"qq": %d
	}`, SessionKey, Config.SNS.QQ.BotQQNumber)
	url = Config.SNS.QQ.APILink + "/bind"
	HttpRequest.PostRequestJson(url, requestBody, []string{})

	//缓存SessionKey
	ioutil.WriteFile("data/SessionKey", []byte(SessionKey), 0664)
	return SessionKey, resp, http_error
}

func GetSessionKey() string {
	bytes, _ := ioutil.ReadFile("data/SessionKey")
	SessionKey := string(bytes)

	return SessionKey
}
