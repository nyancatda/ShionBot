/*
 * @Author: NyanCatda
 * @Date: 2021-10-03 05:25:31
 * @LastEditTime: 2022-01-27 18:17:20
 * @LastEditors: NyanCatda
 * @Description: QQ消息处理
 * @FilePath: \ShionBot\src\MessageProcessing\QQMessageProcessing.go
 */
package MessageProcessing

import (
	"strconv"

	"github.com/nyancatda/ShionBot/Controller/MessageProcessing/Struct"
	"github.com/nyancatda/ShionBot/Controller/MessagePushAPI"
	"github.com/nyancatda/ShionBot/Modular/Command"
	"github.com/nyancatda/ShionBot/Modular/GetWikiInfo"
	"github.com/nyancatda/ShionBot/Utils/Language"
	"github.com/nyancatda/ShionBot/Utils/ReadConfig"
)

var sns_name_qq string = "QQ"

//发送群组消息
func QQsendGroupWikiInfo(json Struct.WebHookJson, UserID string, WikiName string, GroupID string, QueryText string, quoteID string) {
	WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_qq, json, UserID, WikiName, QueryText, "")
	if err != nil {
		WikiLink := ReadConfig.GetWikiLink(sns_name_qq, json, WikiName)
		MessagePushAPI.SendMessage(sns_name_qq, "Group", UserID, GroupID, Error(sns_name_qq, UserID, WikiLink), true, quoteID, "", 0)
		return
	}
	MessagePushAPI.SendMessage(sns_name_qq, "Group", UserID, GroupID, WikiInfo, true, quoteID, "", 0)
}

//发送好友消息
func QQsendFriendWikiInfo(json Struct.WebHookJson, WikiName string, UserID string, QueryText string) {
	WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_qq, json, UserID, WikiName, QueryText, "")
	if err != nil {
		WikiLink := ReadConfig.GetWikiLink(sns_name_qq, json, WikiName)
		MessagePushAPI.SendMessage(sns_name_qq, "Friend", UserID, UserID, Error(sns_name_qq, UserID, WikiLink), false, "", "", 0)
		return
	}
	MessagePushAPI.SendMessage(sns_name_qq, "Friend", UserID, UserID, WikiInfo, false, "", "", 0)
}

//发送临时会话消息
func QQsendTempdWikiInfo(json Struct.WebHookJson, WikiName string, UserID string, GroupID int, QueryText string) {
	WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_qq, json, UserID, WikiName, QueryText, "")
	if err != nil {
		WikiLink := ReadConfig.GetWikiLink(sns_name_qq, json, WikiName)
		MessagePushAPI.SendMessage(sns_name_qq, "Temp", UserID, UserID, Error(sns_name_qq, UserID, WikiLink), false, "", "", GroupID)
		return
	}
	MessagePushAPI.SendMessage(sns_name_qq, "Temp", UserID, UserID, WikiInfo, false, "", "", GroupID)
}

//戳一戳消息处理
func QQNudgeEventMessageProcessing(json Struct.WebHookJson) {
	UserID := json.FromId
	HelpText := Language.Message(sns_name_qq, strconv.Itoa(UserID)).HelpText
	switch json.Subject.Kind {
	case "Group":
		if json.FromId != ReadConfig.GetConfig.SNS.QQ.BotQQNumber && json.Target == ReadConfig.GetConfig.SNS.QQ.BotQQNumber {
			MessagePushAPI.SendNudge(json.FromId, json.Subject.Id, "Group")
			MessagePushAPI.SendMessage(sns_name_qq, "GroupAt", strconv.Itoa(UserID), strconv.Itoa(json.Subject.Id), HelpText, false, "", strconv.Itoa(json.FromId), 0)
		}
	case "Friend":
		go MessagePushAPI.SendMessage(sns_name_qq, "Friend", strconv.Itoa(UserID), strconv.Itoa(UserID), HelpText, false, "", "", 0)
	}
}

//消息处理
func QQMessageProcessing(json Struct.WebHookJson) {
	ChatType := json.Type
	switch ChatType {
	case "GroupMessage":
		//不处理非正常消息
		if len(json.MessageChain) < 2 {
			return
		}
		if json.MessageChain[1].Type == "Plain" {
			text := json.MessageChain[1].Text
			UserID := strconv.Itoa(json.Sender.Id)
			find, Command, CommandData := CommandExtraction(sns_name_qq, json, text)
			if find {
				if Command == "/" {
					QQSettingsMessageProcessing(CommandData, json)
					return
				}
				GroupID := strconv.Itoa(json.Sender.Group.Id)
				quoteID := strconv.Itoa(json.MessageChain[0].Id)
				MessagePushAPI.SendNudge(json.Sender.Id, json.Sender.Group.Id, "Group")
				QQsendGroupWikiInfo(json, UserID, Command, GroupID, CommandData, quoteID)
			}
		}
	case "FriendMessage":
		if len(json.MessageChain) < 2 {
			return
		}
		if json.MessageChain[1].Type == "Plain" {
			text := json.MessageChain[1].Text
			UserID := strconv.Itoa(json.Sender.Id)
			find, Command, CommandData := CommandExtraction(sns_name_qq, json, text)
			if find {
				if Command == "/" {
					QQSettingsMessageProcessing(CommandData, json)
					return
				}
				QQsendFriendWikiInfo(json, Command, UserID, CommandData)
			}
		}
	case "TempMessage":
		if len(json.MessageChain) < 2 {
			return
		}
		if json.MessageChain[1].Type == "Plain" {
			text := json.MessageChain[1].Text
			UserID := strconv.Itoa(json.Sender.Id)
			find, Command, CommandData := CommandExtraction(sns_name_qq, json, text)
			if find {
				if Command == "/" {
					QQSettingsMessageProcessing(CommandData, json)
					return
				}
				GroupID := json.Sender.Group.Id
				QQsendTempdWikiInfo(json, Command, UserID, GroupID, CommandData)
			}
		}
	case "NudgeEvent":
		QQNudgeEventMessageProcessing(json)
	}
}

//设置消息返回
func QQSettingsMessageProcessing(Text string, json Struct.WebHookJson) {
	Message, Bool := Command.Command(sns_name_qq, json, Text)
	if Bool {
		UserID := strconv.Itoa(json.Sender.Id)
		ChatType := json.Type
		switch ChatType {
		case "GroupMessage":
			GroupID := strconv.Itoa(json.Sender.Group.Id)
			quoteID := strconv.Itoa(json.MessageChain[0].Id)
			MessagePushAPI.SendMessage(sns_name_qq, "Group", UserID, GroupID, Message, true, quoteID, "", 0)
		case "FriendMessage":
			UserID := strconv.Itoa(json.Sender.Id)
			MessagePushAPI.SendMessage(sns_name_qq, "Friend", UserID, UserID, Message, false, "", "", 0)
		case "TempMessage":
			UserID := strconv.Itoa(json.Sender.Id)
			GroupID := json.Sender.Group.Id
			MessagePushAPI.SendMessage(sns_name_qq, "Temp", UserID, UserID, Message, false, "", "", GroupID)
		}
	}
}
