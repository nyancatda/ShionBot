package InformationProcessing

import (
	"strconv"

	"xyz.nyan/MediaWiki-Bot/src/MediaWikiAPI"
	"xyz.nyan/MediaWiki-Bot/src/MessageProcessingAPI"
	"xyz.nyan/MediaWiki-Bot/src/Plugin"
	"xyz.nyan/MediaWiki-Bot/src/Struct"
)

func TelegramMessageProcessing(json Struct.WebHookJson) {
	text := json.Message.Text
	find, QueryText, Command := CommandExtraction("Telegram", json, text)
	if find {
		UserID := json.Message.From.Id
		ChatID := json.Message.Chat.Id
		WikiInfo, err := Plugin.GetWikiInfo(UserID, Command, QueryText)
		if err != nil {
			WikiLink := MediaWikiAPI.GetWikiLink(Command)
			go MessageProcessingAPI.SendMessage("Telegram", "Default", ChatID, Error(strconv.Itoa(UserID), WikiLink), false, 0, "", 0)
			return
		}
		switch json.Message.Chat.Type {
		case "private":
			go MessageProcessingAPI.SendMessage("Telegram", "Default", ChatID, WikiInfo, false, 0, "", 0)
		case "supergroup":
			MassageID := json.Message.Message_id
			go MessageProcessingAPI.SendMessage("Telegram", "Group", ChatID, WikiInfo, true, MassageID, "", 0)
		}
	}
}
