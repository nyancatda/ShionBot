package InformationProcessing

import (
	"fmt"
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
		WikiInfo, err := Plugin.GetWikiInfo(UserID, Command, QueryText)
		if err != nil {
			WikiLink := MediaWikiAPI.GetWikiLink(Command)
			MessageProcessingAPI.SendMessage("Telegram", "Default", UserID, Error(strconv.Itoa(UserID), WikiLink), false, 0, "", 0)
			return
		}
		switch json.Message.Chat.Type {
		case "private":
			MessageProcessingAPI.SendMessage("Telegram", "Default", UserID, WikiInfo, false, 0, "", 0)
		case "supergroup":
			MassageID := json.Message.Message_id
			MessageProcessingAPI.SendMessage("Telegram", "GroupAt", UserID, WikiInfo, true, MassageID, "", 0)
		}
	}
	fmt.Println(json.Message.Text)
}
