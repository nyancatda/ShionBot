package LineAPI

import (
	"encoding/json"

	"xyz.nyan/ShionBot/src/MessagePushAPI/SNSAPI"
	"xyz.nyan/ShionBot/src/utils"
)

var sns_name string = "Telegram"

//发送推送消息
//to 消息接收的聊天ID
//messages 消息内容
//notificationDisabled 是否需要静默发送
func SendPushMessage(chat_type string, to string, messages string, notificationDisabled bool) {
	Config := utils.ReadConfig()
	Json := map[string]interface{}{
		"to":                   to,
		"notificationDisabled": notificationDisabled,
	}
	JsonMessages := make([]map[string]string, 1)
	JsonMessages[0] = map[string]string{
		"type": "text",
		"text": messages,
	}
	Json["messages"] = JsonMessages
	JsonBody, _ := json.Marshal(Json)
	requestBody := string(JsonBody)

	url := Config.SNS.Line.BotAPILink + "v2/bot/message/push"
	Header := []string{"Authorization:Bearer " + Config.SNS.Line.ChannelAccessToken}
	utils.PostRequestJosnHeader(url, requestBody, Header)

	SNSAPI.Log(sns_name, chat_type, to, messages)
}

//发送回复消息
//replyToken 消息回复令牌
//messages 消息内容
//notificationDisabled 是否需要静默发送
func SendReplyMessage(chat_type string, replyToken string, messages string, notificationDisabled bool) {
	Config := utils.ReadConfig()
	Json := map[string]interface{}{
		"replyToken":           replyToken,
		"notificationDisabled": notificationDisabled,
	}
	JsonMessages := make([]map[string]string, 1)
	JsonMessages[0] = map[string]string{
		"type": "text",
		"text": messages,
	}
	Json["messages"] = JsonMessages
	JsonBody, _ := json.Marshal(Json)
	requestBody := string(JsonBody)

	url := Config.SNS.Line.BotAPILink + "v2/bot/message/reply"
	Header := []string{"Authorization:Bearer " + Config.SNS.Line.ChannelAccessToken}
	utils.PostRequestJosnHeader(url, requestBody, Header)

	SNSAPI.Log(sns_name, chat_type, replyToken, messages)
}
