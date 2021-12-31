package InformationProcessing

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/nyancatda/ShionBot/src/MediaWikiAPI"
	"github.com/nyancatda/ShionBot/src/MessagePushAPI"
	"github.com/nyancatda/ShionBot/src/Modular/Command"
	"github.com/nyancatda/ShionBot/src/Modular/GetWikiInfo"
	"github.com/nyancatda/ShionBot/src/Struct"
)

var sns_name_kaiheila string = "KaiHeiLa"

func KaiHeiLaWebHookVerifyProcessing(c *gin.Context, json Struct.WebHookJson) {
	Challenge := json.D.Challenge
	JsonData := map[string]interface{}{
		"challenge": Challenge,
	}
	c.JSONP(200, JsonData)
}

func KaiHeiLaMessageProcessing(json Struct.WebHookJson) {
	text := json.D.Content
	find, QueryText, Command := CommandExtraction(sns_name_kaiheila, json, text)
	if find {
		UserID := json.D.Author_id
		ChatType := json.D.Channel_type
		Log(sns_name_kaiheila, ChatType, UserID, text)
		switch ChatType {
		case "PERSON":
			ChatID := json.D.Author_id
			WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_kaiheila, json, UserID, Command, QueryText, "")
			if err != nil {
				WikiLink := MediaWikiAPI.GetWikiLink(sns_name_kaiheila, json, Command)
				MessagePushAPI.SendMessage(sns_name_kaiheila, "Friend", UserID, ChatID, Error(sns_name_kaiheila, UserID, WikiLink), false, "", "", 0)
				return
			}
			MessagePushAPI.SendMessage(sns_name_kaiheila, "Friend", UserID, ChatID, WikiInfo, false, "", "", 0)
		case "GROUP":
			MassageID := json.D.Msg_id
			ChatID := json.D.Target_id
			WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_kaiheila, json, UserID, Command, QueryText, "")
			if err != nil {
				WikiLink := MediaWikiAPI.GetWikiLink(sns_name_kaiheila, json, Command)
				MessagePushAPI.SendMessage(sns_name_kaiheila, "Group", UserID, ChatID, Error(sns_name_kaiheila, UserID, WikiLink), false, "", "", 0)
				return
			}
			MessagePushAPI.SendMessage(sns_name_kaiheila, "Group", UserID, ChatID, WikiInfo, true, MassageID, "", 0)
		}
	}
}

//设置消息返回
func KaiHeiLaSettingsMessageProcessing(json Struct.WebHookJson) {
	text := json.D.Content
	countSplit := strings.SplitN(text, "/", 2)
	Text := countSplit[1]
	Message, Bool := Command.Command(sns_name_kaiheila, json, Text)
	if Bool {
		UserID := json.D.Author_id
		ChatType := json.D.Channel_type
		Log(sns_name_kaiheila, ChatType, UserID, text)
		switch ChatType {
		case "PERSON":
			ChatID := json.D.Author_id
			MessagePushAPI.SendMessage(sns_name_kaiheila, "Friend", UserID, ChatID, Message, false, "", "", 0)
		case "GROUP":
			ChatID := json.D.Target_id
			MassageID := json.D.Msg_id
			MessagePushAPI.SendMessage(sns_name_kaiheila, "Group", UserID, ChatID, Message, true, MassageID, "", 0)
		}
	}
}
