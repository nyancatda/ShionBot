package InformationProcessing

import (
	"math"
	"strconv"
	"strings"

	"xyz.nyan/MediaWiki-Bot/src/MediaWikiAPI"
	"xyz.nyan/MediaWiki-Bot/src/MessageProcessingAPI"
	"xyz.nyan/MediaWiki-Bot/src/Plugin"
	"xyz.nyan/MediaWiki-Bot/src/Struct"
	"xyz.nyan/MediaWiki-Bot/src/utils"
	"xyz.nyan/MediaWiki-Bot/src/utils/Language"
	"xyz.nyan/MediaWiki-Bot/src/utils/Settings"
)

func QQError(UserID string, WikiLink string) string {
	return Language.StringVariable(1, Language.Message(UserID).WikiLinkError, WikiLink, "")
}

//发送群组消息
func QQsendGroupWikiInfo(UserID int, WikiName string, GroupID int, QueryText string, quoteID int) {
	WikiInfo, err := Plugin.GetWikiInfo(UserID, WikiName, QueryText)
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		MessageProcessingAPI.SendGroupMessage("QQ", GroupID, QQError(strconv.Itoa(UserID), WikiLink), true, quoteID)
		return
	}
	MessageProcessingAPI.SendGroupMessage("QQ", GroupID, WikiInfo, true, quoteID)
}

//发送好友消息
func QQsendFriendWikiInfo(WikiName string, UserID int, QueryText string) {
	WikiInfo, err := Plugin.GetWikiInfo(UserID, WikiName, QueryText)
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		MessageProcessingAPI.SendFriendMessage("QQ", UserID, QQError(strconv.Itoa(UserID), WikiLink), false, 0)
		return
	}
	MessageProcessingAPI.SendFriendMessage("QQ", UserID, WikiInfo, false, 0)
}

//发送临时会话消息
func QQsendTempdWikiInfo(WikiName string, UserID int, GroupID int, QueryText string) {
	WikiInfo, err := Plugin.GetWikiInfo(UserID, WikiName, QueryText)
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		MessageProcessingAPI.SendTempMessage("QQ", UserID, GroupID, QQError(strconv.Itoa(UserID), WikiLink), false, 0)
		return
	}
	MessageProcessingAPI.SendTempMessage("QQ", UserID, GroupID, WikiInfo, false, 0)
}

//戳一戳消息处理
func QQNudgeEventMessageProcessing(json Struct.QQWebHook_root) {
	UserID := json.Sender.Id
	HelpText := Language.Message(strconv.Itoa(UserID)).HelpText
	switch json.Subject.Kind {
	case "Group":
		if json.FromId != utils.ReadConfig().QQBot.BotQQNumber && json.Target == utils.ReadConfig().QQBot.BotQQNumber {
			go MessageProcessingAPI.SendNudge(json.FromId, json.Subject.Id, "Group")
			go MessageProcessingAPI.SendGroupAtMessage("QQ", json.Subject.Id, HelpText, json.FromId)
		}
	case "Friend":
		go MessageProcessingAPI.SendFriendMessage("QQ", json.FromId, HelpText, false, 0)
	}
}

//消息处理
func QQMessageProcessing(json Struct.QQWebHook_root) {
	switch json.Type {
	case "GroupMessage":
		if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
			text := json.MessageChain[1].(map[string]interface{})["text"]
			find, QueryText, Command := CommandExtraction("QQ", json, text.(string))
			if find {
				GroupID := json.Sender.Group.Id
				quoteID := int(math.Floor(json.MessageChain[0].(map[string]interface{})["id"].(float64)))
				UserID := json.Sender.Id
				go MessageProcessingAPI.SendNudge(UserID, GroupID, "Group")
				go QQsendGroupWikiInfo(UserID, Command, GroupID, QueryText, quoteID)
			}
		}
	case "FriendMessage":
		if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
			text := json.MessageChain[1].(map[string]interface{})["text"]
			find, QueryText, Command := CommandExtraction("QQ", json, text.(string))
			if find {
				UserID := json.Sender.Id
				go QQsendFriendWikiInfo(Command, UserID, QueryText)
			}
		}
	case "TempMessage":
		if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
			text := json.MessageChain[1].(map[string]interface{})["text"]
			find, QueryText, Command := CommandExtraction("QQ", json, text.(string))
			if find {
				UserID := json.Sender.Id
				GroupID := json.Sender.Group.Id
				go QQsendTempdWikiInfo(Command, UserID, GroupID, QueryText)
			}
		}
	case "NudgeEvent":
		QQNudgeEventMessageProcessing(json)
	}
}

//设置消息返回
func QQSettingsMessageProcessing(json Struct.QQWebHook_root) {
	text := json.MessageChain[1].(map[string]interface{})["text"]
	countSplit := strings.Split(text.(string), "/")
	Text := countSplit[1]
	Message, Bool := Settings.Settings("QQ", json, Text)
	if Bool {
		switch json.Type {
		case "GroupMessage":
			GroupID := json.Sender.Group.Id
			quoteID := int(math.Floor(json.MessageChain[0].(map[string]interface{})["id"].(float64)))
			go MessageProcessingAPI.SendGroupMessage("QQ", GroupID, Message, true, quoteID)
		case "FriendMessage":
			UserID := json.Sender.Id
			go MessageProcessingAPI.SendFriendMessage("QQ", UserID, Message, false, 0)
		case "TempMessage":
			UserID := json.Sender.Id
			GroupID := json.Sender.Group.Id
			go MessageProcessingAPI.SendTempMessage("QQ", UserID, GroupID, Message, false, 0)
		}
	}
}
