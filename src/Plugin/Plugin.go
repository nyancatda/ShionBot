package Plugin

import (
	"strconv"

	"github.com/nyancatda/ShionBot/src/Struct"
)

func GetSNSUserID(SNSName string, Messagejson Struct.WebHookJson) string {
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
	return UserID
}
