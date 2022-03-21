/*
 * @Author: NyanCatda
 * @Date: 2021-11-05 23:42:17
 * @LastEditTime: 2022-03-21 17:34:51
 * @LastEditors: NyanCatda
 * @Description: Line消息处理
 * @FilePath: \ShionBot\Controller\MessageProcessing\LineMessageProcessing.go
 */
package MessageProcessing

import (
	"github.com/nyancatda/ShionBot/Controller/MessageProcessing/Struct"
	"github.com/nyancatda/ShionBot/Controller/MessagePushAPI"
	"github.com/nyancatda/ShionBot/Modular/Command"
	"github.com/nyancatda/ShionBot/Modular/GetWikiInfo"
	"github.com/nyancatda/ShionBot/Utils/ReadConfig"
)

var sns_name_line string = "Line"

func LineMessageProcessing(json Struct.WebHookJson) {
	if len(json.Events) <= 0 {
		return
	}
	text := json.Events[0].Message.Text
	//判断命令是否匹配
	find, Command, CommandData := CommandExtraction(sns_name_line, json, text)
	if find {
		if Command == "/" {
			LineSettingsMessageProcessing(CommandData, json)
			return
		}

		UserID := json.Events[0].Source.UserId
		ChatType := json.Events[0].Source.Type
		switch ChatType {
		case "user":
			WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_line, json, UserID, Command, CommandData, "")
			if err != nil {
				WikiLink := ReadConfig.GetWikiLink(sns_name_line, json, Command)
				MessagePushAPI.SendMessage(sns_name_line, "Default", UserID, UserID, Error(sns_name_line, UserID, WikiLink), false, "", "", 0)
				return
			}
			MessagePushAPI.SendMessage(sns_name_line, "Default", UserID, UserID, WikiInfo, false, "", "", 0)
		case "group":
			GroupId := json.Events[0].Source.GroupId
			QuoteID := json.Events[0].ReplyToken
			WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_line, json, UserID, Command, CommandData, "")
			if err != nil {
				WikiLink := ReadConfig.GetWikiLink(sns_name_line, json, Command)
				MessagePushAPI.SendMessage(sns_name_line, "Group", UserID, GroupId, Error(sns_name_line, UserID, WikiLink), true, QuoteID, "", 0)
				return
			}
			MessagePushAPI.SendMessage(sns_name_line, "Group", UserID, GroupId, WikiInfo, true, QuoteID, "", 0)
		}
	}
}

//设置消息返回
func LineSettingsMessageProcessing(Text string, json Struct.WebHookJson) {
	Message, Bool := Command.Command(sns_name_line, json, Text)
	if Bool {
		ChatType := json.Events[0].Source.Type
		UserID := json.Events[0].Source.UserId
		switch ChatType {
		case "user":
			MessagePushAPI.SendMessage(sns_name_line, "Default", UserID, UserID, Message, false, "", "", 0)
		case "group":
			GroupId := json.Events[0].Source.GroupId
			QuoteID := json.Events[0].ReplyToken
			MessagePushAPI.SendMessage(sns_name_line, "Group", UserID, GroupId, Message, true, QuoteID, "", 0)
		}
	}
}
