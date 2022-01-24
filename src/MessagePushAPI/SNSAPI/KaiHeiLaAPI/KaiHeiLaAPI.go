/*
 * @Author: NyanCatda
 * @Date: 2021-11-15 17:23:29
 * @LastEditTime: 2022-01-24 21:10:50
 * @LastEditors: NyanCatda
 * @Description: KaiHeiLa API
 * @FilePath: \ShionBot\src\MessagePushAPI\SNSAPI\KaiHeiLaAPI\KaiHeiLaAPI.go
 */
package KaiHeiLaAPI

import (
	"encoding/json"
	"net/http"

	"github.com/nyancatda/ShionBot/src/Utils/HttpRequest"
	"github.com/nyancatda/ShionBot/src/Utils/ReadConfig"
)

var (
	APILink = "https://www.kaiheila.cn/"
)

/**
 * @description: 发送私信消息
 * @param {int} Type 消息类型 1 文本类型，2 图片消息，3 视频消息，4 文件消息，9 代表 kmarkdown 消息, 10 代表卡片消息
 * @param {string} target_id 目标用户id
 * @param {string} content 消息内容
 * @param {bool} quote 是否需要回复
 * @param {string} quoteID 回复ID
 * @return {[]byte}
 * @return {*http.Response}
 * @return {error}
 */
func SendDirectMessage(Type int, target_id string, content string, quote bool, quoteID string) ([]byte, *http.Response, error) {
	Config := ReadConfig.GetConfig

	//组成消息
	var Json map[string]interface{}
	if quote {
		Json = map[string]interface{}{
			"type":      Type,
			"target_id": target_id,
			"content":   content,
			"quote":     quoteID,
		}
	} else {
		Json = map[string]interface{}{
			"type":      Type,
			"target_id": target_id,
			"content":   content,
		}
	}
	JsonBody, _ := json.Marshal(Json)

	url := APILink + "api/v3/direct-message/create"
	//请求头添加令牌
	Header := []string{"Authorization:Bot " + Config.SNS.KaiHeiLa.Token}
	Body, HttpResponse, err := HttpRequest.PostRequestJson(url, string(JsonBody), Header)

	return Body, HttpResponse, err
}

/**
 * @description: 发送频道聊天消息
 * @param {int} Type 消息类型 1 文本类型，2 图片消息，3 视频消息，4 文件消息，9 代表 kmarkdown 消息, 10 代表卡片消息
 * @param {string} target_id 目标频道id
 * @param {string} content 消息内容
 * @param {bool} quote 是否需要回复
 * @param {string} quoteID 回复ID
 * @return {[]byte}
 * @return {*http.Response}
 * @return {error}
 */
func SendChannelMessage(Type int, target_id string, content string, quote bool, quoteID string) ([]byte, *http.Response, error) {
	Config := ReadConfig.GetConfig

	//组成消息
	var Json map[string]interface{}
	if quote {
		Json = map[string]interface{}{
			"type":      Type,
			"target_id": target_id,
			"content":   content,
			"quote":     quoteID,
		}
	} else {
		Json = map[string]interface{}{
			"type":      Type,
			"target_id": target_id,
			"content":   content,
		}
	}
	JsonBody, _ := json.Marshal(Json)

	url := APILink + "api/v3/message/create"
	//请求头添加令牌
	Header := []string{"Authorization:Bot " + Config.SNS.KaiHeiLa.Token}
	Body, HttpResponse, err := HttpRequest.PostRequestJson(url, string(JsonBody), Header)

	return Body, HttpResponse, err
}
