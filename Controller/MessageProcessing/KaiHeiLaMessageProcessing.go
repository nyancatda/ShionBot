/*
 * @Author: NyanCatda
 * @Date: 2021-11-15 17:23:29
 * @LastEditTime: 2022-03-07 18:50:11
 * @LastEditors: NyanCatda
 * @Description: KaiHeiLa 消息处理
 * @FilePath: \ShionBot\Controller\MessageProcessing\KaiHeiLaMessageProcessing.go
 */
package MessageProcessing

import (
	"github.com/gin-gonic/gin"

	"github.com/nyancatda/ShionBot/Controller/MessageProcessing/Struct"
	"github.com/nyancatda/ShionBot/Controller/MessagePushAPI"
	"github.com/nyancatda/ShionBot/Modular/Command"
	"github.com/nyancatda/ShionBot/Modular/GetWikiInfo"
	"github.com/nyancatda/ShionBot/Utils/ReadConfig"
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
	//判断命令是否匹配
	find, Command, CommandData := CommandExtraction(sns_name_kaiheila, json, text)
	if find {
		if Command == "/" {
			KaiHeiLaSettingsMessageProcessing(CommandData, json)
			return
		}

		UserID := json.D.Author_id
		ChatType := json.D.Channel_type
		switch ChatType {
		case "PERSON":
			ChatID := json.D.Author_id
			WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_kaiheila, json, UserID, Command, CommandData, "")
			if err != nil {
				WikiLink := ReadConfig.GetWikiLink(sns_name_kaiheila, json, Command)
				MessagePushAPI.SendMessage(sns_name_kaiheila, "Friend", UserID, ChatID, Error(sns_name_kaiheila, UserID, WikiLink), false, "", "", 0)
				return
			}
			MessagePushAPI.SendMessage(sns_name_kaiheila, "Friend", UserID, ChatID, WikiInfo, false, "", "", 0)
		case "GROUP":
			MassageID := json.D.Msg_id
			ChatID := json.D.Target_id
			WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_kaiheila, json, UserID, Command, CommandData, "")
			if err != nil {
				WikiLink := ReadConfig.GetWikiLink(sns_name_kaiheila, json, Command)
				MessagePushAPI.SendMessage(sns_name_kaiheila, "Group", UserID, ChatID, Error(sns_name_kaiheila, UserID, WikiLink), false, "", "", 0)
				return
			}
			MessagePushAPI.SendMessage(sns_name_kaiheila, "Group", UserID, ChatID, WikiInfo, true, MassageID, "", 0)
		}
	}
}

//设置消息返回
func KaiHeiLaSettingsMessageProcessing(Text string, json Struct.WebHookJson) {
	Message, Bool := Command.Command(sns_name_kaiheila, json, Text)
	if Bool {
		UserID := json.D.Author_id
		ChatType := json.D.Channel_type
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
