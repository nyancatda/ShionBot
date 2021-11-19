package InformationProcessing

import (
	"strconv"
	"strings"

	"xyz.nyan/ShionBot/src/MediaWikiAPI"
	"xyz.nyan/ShionBot/src/MessagePushAPI"
	"xyz.nyan/ShionBot/src/Plugin/Command"
	"xyz.nyan/ShionBot/src/Plugin/GetWikiInfo"
	"xyz.nyan/ShionBot/src/Struct"
	"xyz.nyan/ShionBot/src/utils"
	"xyz.nyan/ShionBot/src/utils/Language"
)

var sns_name_qq string = "QQ"

//发送群组消息
func QQsendGroupWikiInfo(json Struct.WebHookJson, UserID string, WikiName string, GroupID string, QueryText string, quoteID string) {
	WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_qq, json, UserID, WikiName, QueryText, "")
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(sns_name_qq, json, WikiName)
		MessagePushAPI.SendMessage(sns_name_qq, "Group", UserID, GroupID, Error(sns_name_qq, UserID, WikiLink), true, quoteID, "", 0)
		return
	}
	MessagePushAPI.SendMessage(sns_name_qq, "Group", UserID, GroupID, WikiInfo, true, quoteID, "", 0)
}

//发送好友消息
func QQsendFriendWikiInfo(json Struct.WebHookJson, WikiName string, UserID string, QueryText string) {
	WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_qq, json, UserID, WikiName, QueryText, "")
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(sns_name_qq, json, WikiName)
		MessagePushAPI.SendMessage(sns_name_qq, "Friend", UserID, UserID, Error(sns_name_qq, UserID, WikiLink), false, "", "", 0)
		return
	}
	MessagePushAPI.SendMessage(sns_name_qq, "Friend", UserID, UserID, WikiInfo, false, "", "", 0)
}

//发送临时会话消息
func QQsendTempdWikiInfo(json Struct.WebHookJson, WikiName string, UserID string, GroupID int, QueryText string) {
	WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_qq, json, UserID, WikiName, QueryText, "")
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(sns_name_qq, json, WikiName)
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
		if json.FromId != utils.ReadConfig().SNS.QQ.BotQQNumber && json.Target == utils.ReadConfig().SNS.QQ.BotQQNumber {
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
			find, QueryText, Command := CommandExtraction(sns_name_qq, json, text)
			if find {
				GroupID := strconv.Itoa(json.Sender.Group.Id)
				quoteID := strconv.Itoa(json.MessageChain[0].Id)
				Log(sns_name_qq, ChatType, UserID, text)
				MessagePushAPI.SendNudge(json.Sender.Id, json.Sender.Group.Id, "Group")
				QQsendGroupWikiInfo(json, UserID, Command, GroupID, QueryText, quoteID)
			}
		}
	case "FriendMessage":
		if len(json.MessageChain) < 2 {
			return
		}
		if json.MessageChain[1].Type == "Plain" {
			text := json.MessageChain[1].Text
			UserID := strconv.Itoa(json.Sender.Id)
			find, QueryText, Command := CommandExtraction(sns_name_qq, json, text)
			if find {
				Log(sns_name_qq, ChatType, UserID, text)
				QQsendFriendWikiInfo(json, Command, UserID, QueryText)
			}
		}
	case "TempMessage":
		if len(json.MessageChain) < 2 {
			return
		}
		if json.MessageChain[1].Type == "Plain" {
			text := json.MessageChain[1].Text
			UserID := strconv.Itoa(json.Sender.Id)
			find, QueryText, Command := CommandExtraction(sns_name_qq, json, text)
			if find {
				GroupID := json.Sender.Group.Id
				Log(sns_name_qq, ChatType, UserID, text)
				QQsendTempdWikiInfo(json, Command, UserID, GroupID, QueryText)
			}
		}
	case "NudgeEvent":
		text := Language.DefaultLanguageMessage().Nudge
		UserID := strconv.Itoa(json.FromId)
		Log(sns_name_qq, ChatType, UserID, text)
		QQNudgeEventMessageProcessing(json)
	}
}

//设置消息返回
func QQSettingsMessageProcessing(json Struct.WebHookJson) {
	text := json.MessageChain[1].Text
	countSplit := strings.SplitN(text, "/", 2)
	Text := countSplit[1]
	Message, Bool := Command.Command(sns_name_qq, json, Text)
	if Bool {
		UserID := strconv.Itoa(json.Sender.Id)
		ChatType := json.Type
		Log(sns_name_qq, ChatType, UserID, text)
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
