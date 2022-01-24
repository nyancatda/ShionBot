package TelegramAPI

import (
	"fmt"
	"strconv"

	"github.com/nyancatda/ShionBot/src/MessagePushAPI/SNSAPI"
	"github.com/nyancatda/ShionBot/src/Utils"
)

var sns_name string = "Telegram"

//发送消息
//chat_id 聊天ID
//text 需要发送的信息
//disable_web_page_preview 是否需要禁用链接预览
//disable_notification 是否需要静默发送
//reply_to_message_id 需要回复消息的ID
//allow_sending_without_reply 没有找到需要回复的消息时，是否发送
func SendMessage(chat_type string, chat_id int, text string, disable_web_page_preview bool, disable_notification bool, reply_to_message_id int, allow_sending_without_reply bool) {
	Config := Utils.GetConfig
	requestBody := fmt.Sprintf(`{
		"chat_id": %d,
		"text": "%s",
		"disable_web_page_preview": %t,
		"disable_notification": %t,
		"reply_to_message_id": %d,
		"allow_sending_without_reply": %t
	  }`, chat_id, text, disable_web_page_preview, disable_notification, reply_to_message_id, allow_sending_without_reply)

	url := Config.SNS.Telegram.BotAPILink + "bot" + Config.SNS.Telegram.Token + "/sendMessage"
	Utils.PostRequestJosn(url, requestBody)

	SNSAPI.Log(sns_name, chat_type, strconv.Itoa(chat_id), text)
}
