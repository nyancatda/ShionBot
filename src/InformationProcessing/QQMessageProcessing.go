package InformationProcessing

import (
	"math"
	"strconv"
	"strings"

	"xyz.nyan/MediaWiki-Bot/src/MediaWikiAPI"
	"xyz.nyan/MediaWiki-Bot/src/MessagePushAPI"
	"xyz.nyan/MediaWiki-Bot/src/Plugin"
	"xyz.nyan/MediaWiki-Bot/src/Plugin/Command"
	"xyz.nyan/MediaWiki-Bot/src/Struct"
	"xyz.nyan/MediaWiki-Bot/src/utils"
	"xyz.nyan/MediaWiki-Bot/src/utils/Language"
)

var sns_name_qq string = "QQ"

//发送群组消息
func QQsendGroupWikiInfo(UserID int, WikiName string, GroupID int, QueryText string, quoteID int) {
	WikiInfo, err := Plugin.GetWikiInfo(sns_name_qq, UserID, WikiName, QueryText)
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		MessagePushAPI.SendMessage(sns_name_qq, "Group", GroupID, Error(sns_name_qq, strconv.Itoa(UserID), WikiLink), true, quoteID, "", 0)
		return
	}
	MessagePushAPI.SendMessage(sns_name_qq, "Group", GroupID, WikiInfo, true, quoteID, "", 0)
}

//发送好友消息
func QQsendFriendWikiInfo(WikiName string, UserID int, QueryText string) {
	WikiInfo, err := Plugin.GetWikiInfo(sns_name_qq, UserID, WikiName, QueryText)
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		MessagePushAPI.SendMessage(sns_name_qq, "Friend", UserID, Error(sns_name_qq, strconv.Itoa(UserID), WikiLink), false, 0, "", 0)
		return
	}
	MessagePushAPI.SendMessage(sns_name_qq, "Friend", UserID, WikiInfo, false, 0, "", 0)
}

//发送临时会话消息
func QQsendTempdWikiInfo(WikiName string, UserID int, GroupID int, QueryText string) {
	WikiInfo, err := Plugin.GetWikiInfo(sns_name_qq, UserID, WikiName, QueryText)
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		MessagePushAPI.SendMessage(sns_name_qq, "Temp", UserID, Error(sns_name_qq, strconv.Itoa(UserID), WikiLink), false, 0, "", GroupID)
		return
	}
	MessagePushAPI.SendMessage(sns_name_qq, "Temp", UserID, WikiInfo, false, 0, "", GroupID)
}

//戳一戳消息处理
func QQNudgeEventMessageProcessing(json Struct.WebHookJson) {
	UserID := json.Sender.Id
	HelpText := Language.Message(sns_name_qq, strconv.Itoa(UserID)).HelpText
	switch json.Subject.Kind {
	case "Group":
		if json.FromId != utils.ReadConfig().SNS.QQ.BotQQNumber && json.Target == utils.ReadConfig().SNS.QQ.BotQQNumber {
			MessagePushAPI.SendNudge(json.FromId, json.Subject.Id, "Group")
			MessagePushAPI.SendMessage(sns_name_qq, "GroupAt", json.Subject.Id, HelpText, false, 0, strconv.Itoa(json.FromId), 0)
		}
	case "Friend":
		go MessagePushAPI.SendMessage(sns_name_qq, "Friend", json.FromId, HelpText, false, 0, "", 0)
	}
}

//消息处理
func QQMessageProcessing(json Struct.WebHookJson) {
	switch json.Type {
	case "GroupMessage":
		if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
			text := json.MessageChain[1].(map[string]interface{})["text"]
			find, QueryText, Command := CommandExtraction(sns_name_qq, json, text.(string))
			if find {
				GroupID := json.Sender.Group.Id
				quoteID := int(math.Floor(json.MessageChain[0].(map[string]interface{})["id"].(float64)))
				UserID := json.Sender.Id
				MessagePushAPI.SendNudge(UserID, GroupID, "Group")
				QQsendGroupWikiInfo(UserID, Command, GroupID, QueryText, quoteID)
			}
		}
	case "FriendMessage":
		if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
			text := json.MessageChain[1].(map[string]interface{})["text"]
			find, QueryText, Command := CommandExtraction(sns_name_qq, json, text.(string))
			if find {
				UserID := json.Sender.Id
				QQsendFriendWikiInfo(Command, UserID, QueryText)
			}
		}
	case "TempMessage":
		if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
			text := json.MessageChain[1].(map[string]interface{})["text"]
			find, QueryText, Command := CommandExtraction(sns_name_qq, json, text.(string))
			if find {
				UserID := json.Sender.Id
				GroupID := json.Sender.Group.Id
				QQsendTempdWikiInfo(Command, UserID, GroupID, QueryText)
			}
		}
	case "NudgeEvent":
		QQNudgeEventMessageProcessing(json)
	}
}

//设置消息返回
func QQSettingsMessageProcessing(json Struct.WebHookJson) {
	text := json.MessageChain[1].(map[string]interface{})["text"]
	countSplit := strings.Split(text.(string), "/")
	Text := countSplit[1]
	Message, Bool := Command.Command(sns_name_qq, json, Text)
	if Bool {
		switch json.Type {
		case "GroupMessage":
			GroupID := json.Sender.Group.Id
			quoteID := int(math.Floor(json.MessageChain[0].(map[string]interface{})["id"].(float64)))
			MessagePushAPI.SendMessage(sns_name_qq, "Group", GroupID, Message, true, quoteID, "", 0)
		case "FriendMessage":
			UserID := json.Sender.Id
			MessagePushAPI.SendMessage(sns_name_qq, "Friend", UserID, Message, false, 0, "", 0)
		case "TempMessage":
			UserID := json.Sender.Id
			GroupID := json.Sender.Group.Id
			MessagePushAPI.SendMessage(sns_name_qq, "Temp", UserID, Message, false, 0, "", GroupID)
		}
	}
}
