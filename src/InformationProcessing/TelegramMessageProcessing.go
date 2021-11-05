package InformationProcessing

import (
	"strconv"
	"strings"

	"xyz.nyan/MediaWiki-Bot/src/MediaWikiAPI"
	"xyz.nyan/MediaWiki-Bot/src/MessagePushAPI"
	"xyz.nyan/MediaWiki-Bot/src/Plugin"
	"xyz.nyan/MediaWiki-Bot/src/Plugin/Command"
	"xyz.nyan/MediaWiki-Bot/src/Struct"
)

var sns_name_telegram string = "Telegram"

func TelegramMessageProcessing(json Struct.WebHookJson) {
	text := json.Message.Text
	find, QueryText, Command := CommandExtraction(sns_name_telegram, json, text)
	if find {
		UserID := json.Message.From.Id
		ChatID := json.Message.Chat.Id
		WikiInfo, err := Plugin.GetWikiInfo(sns_name_telegram,UserID, Command, QueryText)
		if err != nil {
			WikiLink := MediaWikiAPI.GetWikiLink(Command)
			go MessagePushAPI.SendMessage(sns_name_telegram, "Default", ChatID, Error(sns_name_telegram,strconv.Itoa(UserID), WikiLink), false, 0, "", 0)
			return
		}
		switch json.Message.Chat.Type {
		case "private":
			go MessagePushAPI.SendMessage(sns_name_telegram, "Default", ChatID, WikiInfo, false, 0, "", 0)
		case "supergroup":
			MassageID := json.Message.Message_id
			go MessagePushAPI.SendMessage(sns_name_telegram, "Group", ChatID, WikiInfo, true, MassageID, "", 0)
		}
	}
}

//设置消息返回
func TelegramSettingsMessageProcessing(json Struct.WebHookJson) {
	text := json.Message.Text
	countSplit := strings.Split(text, "/")
	Text := countSplit[1]
	Message, Bool := Command.Command(sns_name_telegram, json, Text)
	if Bool {
		ChatID := json.Message.Chat.Id
		switch json.Message.Chat.Type {
		case "private":
			MessagePushAPI.SendMessage(sns_name_telegram, "Default", ChatID, Message, false, 0, "", 0)
		case "supergroup":
			MassageID := json.Message.Message_id
			MessagePushAPI.SendMessage(sns_name_telegram, "Group", ChatID, Message, true, MassageID, "", 0)
		}
	}
}
