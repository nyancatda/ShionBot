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
		UserID := strconv.Itoa(json.Sender.Id)
		ChatID := strconv.Itoa(json.Message.Chat.Id)
		WikiInfo, err := Plugin.GetWikiInfo(sns_name_telegram, UserID, Command, QueryText, "")
		if err != nil {
			WikiLink := MediaWikiAPI.GetWikiLink(Command)
			MessagePushAPI.SendMessage(sns_name_telegram, "Default", ChatID, Error(sns_name_telegram, UserID, WikiLink), false, "", "", 0)
			return
		}
		switch json.Message.Chat.Type {
		case "private":
			MessagePushAPI.SendMessage(sns_name_telegram, "Default", ChatID, WikiInfo, false, "", "", 0)
		case "supergroup":
			MassageID := strconv.Itoa(json.Message.Message_id)
			MessagePushAPI.SendMessage(sns_name_telegram, "Group", ChatID, WikiInfo, true, MassageID, "", 0)
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
		ChatID := strconv.Itoa(json.Message.Chat.Id)
		switch json.Message.Chat.Type {
		case "private":
			MessagePushAPI.SendMessage(sns_name_telegram, "Default", ChatID, Message, false, "", "", 0)
		case "supergroup":
			MassageID := strconv.Itoa(json.Message.Message_id)
			MessagePushAPI.SendMessage(sns_name_telegram, "Group", ChatID, Message, true, MassageID, "", 0)
		}
	}
}
