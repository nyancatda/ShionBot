package Command

import (
	"strconv"
	"strings"

	"xyz.nyan/MediaWiki-Bot/src/Struct"
)

func Command(SNSName string, Messagejson Struct.WebHookJson, CommandText string) (string, bool) {
	var UserID string
	switch SNSName {
	case "QQ":
		UserID = strconv.Itoa(Messagejson.Sender.Id)
	case "Telegram":
		UserID = strconv.Itoa(Messagejson.Message.From.Id)
	case "Line":
		UserID = Messagejson.Events[0].Source.UserId
	}

	if find := strings.Contains(CommandText, "language"); find {
		return LanguageSettings(SNSName, UserID, CommandText)
	}
	if find := strings.Contains(CommandText, "help"); find {
		return Help(SNSName, UserID)
	}
	return "", false
}
