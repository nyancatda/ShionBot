package Command

import (
	"strconv"
	"strings"

	"xyz.nyan/MediaWiki-Bot/src/Struct"
)

func Command(SNSName string, Messagejson Struct.WebHookJson, CommandText string) (string, bool) {
	if find := strings.Contains(CommandText, "language"); find {
		countSplit := strings.SplitN(CommandText, " ", 2)
		Language := countSplit[1]
		var UserID int
		switch SNSName {
		case "QQ":
			UserID = Messagejson.Sender.Id
		case "Telegram":
			UserID = Messagejson.Message.From.Id
		}
		return LanguageSettings(SNSName, strconv.Itoa(UserID), Language)
	}
	if find := strings.Contains(CommandText, "help"); find {
		var UserID int
		switch SNSName {
		case "QQ":
			UserID = Messagejson.Sender.Id
		case "Telegram":
			UserID = Messagejson.Message.From.Id
		}
		return Help(SNSName, strconv.Itoa(UserID))
	}
	return "", false
}
