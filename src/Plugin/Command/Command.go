package Command

import (
	"strconv"
	"strings"

	"xyz.nyan/ShionBot/src/Struct"
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
	case "KaiHeiLa":
		UserID = Messagejson.D.Author_id
	}

	if find := strings.Contains(CommandText, "language"); find {
		return LanguageSettings(SNSName, UserID, CommandText)
	}
	if find := strings.Contains(CommandText, "help"); find {
		return Help(SNSName, UserID)
	}
	if find := strings.Contains(CommandText, "addwiki"); find {
		return AddWiki(SNSName, UserID, CommandText)
	}
	return "", false
}
