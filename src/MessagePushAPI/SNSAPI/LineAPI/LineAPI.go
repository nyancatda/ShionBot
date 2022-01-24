/*
 * @Author: NyanCatda
 * @Date: 2021-11-05 23:42:17
 * @LastEditTime: 2022-01-24 20:20:23
 * @LastEditors: NyanCatda
 * @Description: Line API
 * @FilePath: \ShionBot\src\MessagePushAPI\SNSAPI\LineAPI\LineAPI.go
 */
package LineAPI

import (
	"encoding/json"
	"net/http"

	"github.com/nyancatda/ShionBot/src/MessagePushAPI/SNSAPI"
	"github.com/nyancatda/ShionBot/src/Utils/HttpRequest"
	"github.com/nyancatda/ShionBot/src/Utils/ReadConfig"
)

var (
	SNSName = "Line"
)

/**
 * @description: 发送消息
 * @param {string} chat_type 聊天类型
 * @param {string} to 消息接收的聊天ID
 * @param {string} messages 消息内容
 * @param {bool} notificationDisabled 是否需要静默发送
 * @return {[]byte}
 * @return {*http.Response}
 * @return {error}
 */
func SendPushMessage(chat_type string, to string, messages string, notificationDisabled bool) ([]byte, *http.Response, error) {
	Config := ReadConfig.GetConfig

	//组成消息链
	MessageChain := make([]map[string]string, 1)
	MessageChain[0] = map[string]string{
		"type": "text",
		"text": messages,
	}
	MessagesJson := map[string]interface{}{
		"to":                   to,
		"notificationDisabled": notificationDisabled,
		"messages":             MessageChain,
	}

	JsonBody, _ := json.Marshal(MessagesJson)
	requestBody := string(JsonBody)

	url := Config.SNS.Line.BotAPILink + "v2/bot/message/push"
	//请求头添加令牌
	Header := []string{"Authorization:Bearer " + Config.SNS.Line.ChannelAccessToken}
	Body, HttpResponse, err := HttpRequest.PostRequestJson(url, requestBody, Header)

	//没有遇到错误则写入日志
	if err != nil {
		SNSAPI.Log(SNSName, chat_type, to, messages)
	}

	return Body, HttpResponse, err
}

/**
 * @description: 发送回复消息
 * @param {string} chat_type 聊天类型
 * @param {string} replyToken 消息回复令牌
 * @param {string} messages 消息内容
 * @param {bool} notificationDisabled 是否需要静默发送
 * @return {[]byte}
 * @return {*http.Response}
 * @return {error}
 */
func SendReplyMessage(chat_type string, replyToken string, messages string, notificationDisabled bool) ([]byte, *http.Response, error) {
	Config := ReadConfig.GetConfig

	//组成消息链
	MessageChain := make([]map[string]string, 1)
	MessageChain[0] = map[string]string{
		"type": "text",
		"text": messages,
	}
	Json := map[string]interface{}{
		"replyToken":           replyToken,
		"notificationDisabled": notificationDisabled,
		"messages":             MessageChain,
	}

	JsonBody, _ := json.Marshal(Json)
	requestBody := string(JsonBody)

	url := Config.SNS.Line.BotAPILink + "v2/bot/message/reply"
	//请求头添加令牌
	Header := []string{"Authorization:Bearer " + Config.SNS.Line.ChannelAccessToken}
	Body, HttpResponse, err := HttpRequest.PostRequestJson(url, requestBody, Header)

	//没有遇到错误则写入日志
	if err != nil {
		SNSAPI.Log(SNSName, chat_type, replyToken, messages)
	}

	return Body, HttpResponse, err
}
