/*
 * @Author: NyanCatda
 * @Date: 2021-11-15 17:23:29
 * @LastEditTime: 2022-01-24 19:39:33
 * @LastEditors: NyanCatda
 * @Description: KaiHeiLa API
 * @FilePath: \ShionBot\src\MessagePushAPI\SNSAPI\KaiHeiLaAPI\KaiHeiLaAPI.go
 */
package KaiHeiLaAPI

import (
	"encoding/json"
	"fmt"

	"github.com/nyancatda/ShionBot/src/MessagePushAPI/SNSAPI"
	"github.com/nyancatda/ShionBot/src/Utils"
	"github.com/nyancatda/ShionBot/src/Utils/ReadConfig"
)

var sns_name string = "KaiHeiLa"

var APILink string = "https://www.kaiheila.cn/"

//发送私信聊天消息
//Type 消息类型 1 文本类型，2 图片消息，3 视频消息，4 文件消息，9 代表 kmarkdown 消息, 10 代表卡片消息
//target_id 目标用户id
//content 消息内容
//quote 是否需要回复
//quoteID 回复ID
func SendDirectMessage(chat_type string, Type int, target_id string, content string, quote bool, quoteID string) {
	Config := ReadConfig.GetConfig
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
	requestBody := string(JsonBody)

	url := APILink + "api/v3/direct-message/create"
	Header := []string{"Authorization:Bot " + Config.SNS.KaiHeiLa.Token}
	Body, _, _ := Utils.PostRequestJosnHeader(url, requestBody, Header)
	fmt.Println(string(Body))

	SNSAPI.Log(sns_name, chat_type, target_id, content)
}

//发送频道聊天消息
//Type 消息类型 1 文本类型，2 图片消息，3 视频消息，4 文件消息，9 代表 kmarkdown 消息, 10 代表卡片消息
//target_id 目标频道id
//content 消息内容
//quote 是否需要回复
//quoteID 回复ID
func SendChannelMessage(chat_type string, Type int, target_id string, content string, quote bool, quoteID string) {
	Config := ReadConfig.GetConfig
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
	requestBody := string(JsonBody)

	url := APILink + "api/v3/message/create"
	Header := []string{"Authorization:Bot " + Config.SNS.KaiHeiLa.Token}
	Body, _, _ := Utils.PostRequestJosnHeader(url, requestBody, Header)
	fmt.Println(string(Body))

	SNSAPI.Log(sns_name, chat_type, target_id, content)
}
