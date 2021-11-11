package InformationProcessing

import (
	"github.com/gin-gonic/gin"

	"xyz.nyan/MediaWiki-Bot/src/MediaWikiAPI"
	"xyz.nyan/MediaWiki-Bot/src/MessagePushAPI"
	"xyz.nyan/MediaWiki-Bot/src/Plugin"
	"xyz.nyan/MediaWiki-Bot/src/Struct"
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
		switch json.D.Channel_type {
		case "PERSON":
			ChatID := json.D.Author_id
			WikiInfo, err := Plugin.GetWikiInfo(sns_name_kaiheila, UserID, Command, QueryText, "")
			if err != nil {
				WikiLink := MediaWikiAPI.GetWikiLink(Command)
				MessagePushAPI.SendMessage(sns_name_kaiheila, "Default", ChatID, Error(sns_name_kaiheila, UserID, WikiLink), false, "", "", 0)
				return
			}
			MessagePushAPI.SendMessage(sns_name_kaiheila, "Friend", ChatID, WikiInfo, false, "", "", 0)
		case "GROUP":
			MassageID := json.D.Msg_id
			ChatID := json.D.Target_id
			WikiInfo, err := Plugin.GetWikiInfo(sns_name_kaiheila, UserID, Command, QueryText, "")
			if err != nil {
				WikiLink := MediaWikiAPI.GetWikiLink(Command)
				MessagePushAPI.SendMessage(sns_name_kaiheila, "Default", ChatID, Error(sns_name_kaiheila, UserID, WikiLink), false, "", "", 0)
				return
			}
			MessagePushAPI.SendMessage(sns_name_kaiheila, "Group", ChatID, WikiInfo, true, MassageID, "", 0)
		}
	}
}
