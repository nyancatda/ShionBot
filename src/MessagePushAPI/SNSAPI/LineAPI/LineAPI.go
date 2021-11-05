package LineAPI

import (
	"fmt"
	"xyz.nyan/MediaWiki-Bot/src/MessagePushAPI/SNSAPI"
	"xyz.nyan/MediaWiki-Bot/src/utils"
)

var sns_name string = "Telegram"

//发送推送消息
//to 消息接收的聊天ID
//messages 消息内容
//notificationDisabled 是否需要静默发送
func SendPushMessage(chat_type string, to string, messages string, notificationDisabled bool) {
	Config := utils.ReadConfig()
	requestBody := fmt.Sprintf(`{
		"to": "%s",
		"messages": [
		  {
			"type": "text",
			"text": "%s"
		  }
		],
		"notificationDisabled": %t
	  }`, to, messages, notificationDisabled)
	url := "https://api.line.me/v2/bot/message/push"
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
	requestBody := fmt.Sprintf(`{
		"replyToken": "%s",
		"messages": [
		  {
			"type": "text",
			"text": "%s"
		  }
		],
		"notificationDisabled": %t
	  }`, replyToken, messages, notificationDisabled)
	url := "https://api.line.me/v2/bot/message/reply"
	Header := []string{"Authorization:Bearer " + Config.SNS.Line.ChannelAccessToken}
	utils.PostRequestJosnHeader(url, requestBody, Header)

	SNSAPI.Log(sns_name, chat_type, replyToken, messages)
}
