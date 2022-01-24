/*
 * @Author: NyanCatda
 * @Date: 2021-11-05 13:51:15
 * @LastEditTime: 2022-01-24 20:11:08
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \ShionBot\src\MessagePushAPI\SNSAPI\TelegramAPI\TelegramAPI.go
 */
package TelegramAPI

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nyancatda/ShionBot/src/MessagePushAPI/SNSAPI"
	"github.com/nyancatda/ShionBot/src/Utils/HttpRequest"
	"github.com/nyancatda/ShionBot/src/Utils/ReadConfig"
)

var (
	SNSName = "Telegram"
)

//发送消息
//chat_id 聊天ID
//text 需要发送的信息
//disable_web_page_preview 是否需要禁用链接预览
//disable_notification 是否需要静默发送
//reply_to_message_id 需要回复消息的ID
//allow_sending_without_reply 没有找到需要回复的消息时，是否发送
func SendMessage(chat_type string, chat_id int, text string, disable_web_page_preview bool, disable_notification bool, reply_to_message_id int, allow_sending_without_reply bool) ([]byte, *http.Response, error) {
	Config := ReadConfig.GetConfig

	//组成消息Json
	MessageBody := map[string]interface{}{
		"chat_id":                     chat_id,
		"text":                        text,
		"disable_web_page_preview":    disable_web_page_preview,
		"disable_notification":        disable_notification,
		"reply_to_message_id":         reply_to_message_id,
		"allow_sending_without_reply": allow_sending_without_reply,
	}
	requestBody, _ := json.Marshal(MessageBody)

	url := Config.SNS.Telegram.BotAPILink + "bot" + Config.SNS.Telegram.Token + "/sendMessage"
	Body, HttpResponse, err := HttpRequest.PostRequestJson(url, string(requestBody), []string{})

	//没有遇到错误则写入日志
	if err != nil {
		SNSAPI.Log(SNSName, chat_type, strconv.Itoa(chat_id), text)
	}

	return Body, HttpResponse, err
}
