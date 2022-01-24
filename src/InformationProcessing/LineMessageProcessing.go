/*
 * @Author: NyanCatda
 * @Date: 2021-11-05 23:42:17
 * @LastEditTime: 2022-01-24 18:09:14
 * @LastEditors: NyanCatda
 * @Description: Line消息处理
 * @FilePath: \ShionBot\src\InformationProcessing\LineMessageProcessing.go
 */
package InformationProcessing

import (
	"strings"

	"github.com/nyancatda/ShionBot/src/MessagePushAPI"
	"github.com/nyancatda/ShionBot/src/Modular/Command"
	"github.com/nyancatda/ShionBot/src/Modular/GetWikiInfo"
	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/utils"
)

var sns_name_line string = "Line"

func LineMessageProcessing(json Struct.WebHookJson) {
	text := json.Events[0].Message.Text
	find, QueryText, Command := CommandExtraction(sns_name_line, json, text)
	if find {
		UserID := json.Events[0].Source.UserId
		ChatType := json.Events[0].Source.Type
		Log(sns_name_line, ChatType, UserID, text)
		switch ChatType {
		case "user":
			WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_line, json, UserID, Command, QueryText, "")
			if err != nil {
				WikiLink := utils.GetWikiLink(sns_name_line, json, Command)
				MessagePushAPI.SendMessage(sns_name_line, "Default", UserID, UserID, Error(sns_name_line, UserID, WikiLink), false, "", "", 0)
				return
			}
			MessagePushAPI.SendMessage(sns_name_line, "Default", UserID, UserID, WikiInfo, false, "", "", 0)
		case "group":
			GroupId := json.Events[0].Source.GroupId
			QuoteID := json.Events[0].ReplyToken
			WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_line, json, UserID, Command, QueryText, "")
			if err != nil {
				WikiLink := utils.GetWikiLink(sns_name_line, json, Command)
				MessagePushAPI.SendMessage(sns_name_line, "Group", UserID, GroupId, Error(sns_name_line, UserID, WikiLink), true, QuoteID, "", 0)
				return
			}
			MessagePushAPI.SendMessage(sns_name_line, "Group", UserID, GroupId, WikiInfo, true, QuoteID, "", 0)
		}
	}
}

//设置消息返回
func LineSettingsMessageProcessing(json Struct.WebHookJson) {
	text := json.Events[0].Message.Text
	countSplit := strings.SplitN(text, "/", 2)
	Text := countSplit[1]
	Message, Bool := Command.Command(sns_name_line, json, Text)
	if Bool {
		ChatType := json.Events[0].Source.Type
		UserID := json.Events[0].Source.UserId
		Log(sns_name_line, ChatType, UserID, text)
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
