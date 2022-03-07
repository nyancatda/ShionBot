/*
 * @Author: NyanCatda
 * @Date: 2021-11-05 13:51:15
 * @LastEditTime: 2022-01-24 21:19:31
 * @LastEditors: NyanCatda
 * @Description: Session处理API
 * @FilePath: \ShionBot\src\MessagePushAPI\SNSAPI\QQAPI\SessionKey.go
 */
package QQAPI

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nyancatda/ShionBot/Utils/HttpRequest"
	"github.com/nyancatda/ShionBot/Utils/ReadConfig"
)

type verifyJson struct {
	Code    int    `json:"code"`
	Session string `json:"session"`
}

/**
 * @description: 创建一个SessionKey
 * @param {*}
 * @return {string} 新的SessionKey
 * @return {[]byte}
 * @return {*http.Response}
 * @return {error}
 */
func CreateSessionKey() (string, []byte, *http.Response, error) {
	//释放旧的SessionKey
	bytes, _ := ioutil.ReadFile("data/SessionKey")
	OldSessionKey := string(bytes)
	Config := ReadConfig.GetConfig

	//组成请求消息
	MessageBody := map[string]interface{}{
		"verifyKey": OldSessionKey,
		"qq":        Config.SNS.QQ.BotQQNumber,
	}

	MessageBodyJson, _ := json.Marshal(MessageBody)
	url := Config.SNS.QQ.APILink + "/release"
	Body, HttpResponse, err := HttpRequest.PostRequestJson(url, string(MessageBodyJson), []string{})
	if err != nil {
		return "", Body, HttpResponse, err
	}

	//生成一个新的SessionKey
	var SessionKey string
	Config = ReadConfig.GetConfig

	//组成请求消息
	MessageBody = map[string]interface{}{
		"verifyKey": Config.SNS.QQ.VerifyKey,
	}

	MessageBodyJson, _ = json.Marshal(MessageBody)
	url = Config.SNS.QQ.APILink + "/verify"
	Body, HttpResponse, err = HttpRequest.PostRequestJson(url, string(MessageBodyJson), []string{})
	if err != nil {
		return "", Body, HttpResponse, err
	}

	var config verifyJson
	json.Unmarshal([]byte(Body), &config)
	SessionKey = config.Session

	//绑定Key与QQ
	//组成请求消息
	MessageBody = map[string]interface{}{
		"sessionKey": SessionKey,
		"qq":         Config.SNS.QQ.BotQQNumber,
	}
	MessageBodyJson, _ = json.Marshal(MessageBody)
	url = Config.SNS.QQ.APILink + "/bind"
	Body, HttpResponse, err = HttpRequest.PostRequestJson(url, string(MessageBodyJson), []string{})

	//缓存SessionKey
	ioutil.WriteFile("data/SessionKey", []byte(SessionKey), 0664)
	return SessionKey, Body, HttpResponse, err
}

/**
 * @description: 获取缓存本地的SessionKey
 * @param {*}
 * @return {string}
 */
func GetSessionKey() string {
	bytes, err := ioutil.ReadFile("data/SessionKey")

	//如果读取不到则重新生成
	if err != nil {
		SessionKey, _, _, _ := CreateSessionKey()
		return SessionKey
	}

	SessionKey := string(bytes)

	return SessionKey
}
