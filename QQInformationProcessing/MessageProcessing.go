package QQInformationProcessing

import (
	"math"
	"strings"

	"xyz.nyan/MediaWiki-Bot/MediaWikiAPI"
	"xyz.nyan/MediaWiki-Bot/Plugin"
	"xyz.nyan/MediaWiki-Bot/Struct"
	"xyz.nyan/MediaWiki-Bot/utils"
	"xyz.nyan/MediaWiki-Bot/utils/Language"
	"xyz.nyan/MediaWiki-Bot/utils/Settings"
)

func Error(WikiLink string) string {
	return Language.StringVariable(1, Language.Message().WikiLinkError, WikiLink, "")
}

//命令处理，判断命令是否匹配，匹配则输出命令和命令参数
func CommandExtraction(json Struct.QQWebHook_root, text string) (bool, string, string) {
	if find := strings.Contains(text, ":"); find {
		Config := utils.ReadConfig()
		var ConfigWikiName string
		for one := range Config.Wiki.([]interface{}) {
			ConfigWikiName = Config.Wiki.([]interface{})[one].(map[interface{}]interface{})["WikiName"].(string)
			if find := strings.Contains(text, ConfigWikiName); find {
				countSplit := strings.SplitN(text, ":", 2)
				Command := countSplit[0]
				Text := countSplit[1]
				return true, Text, Command
			}
		}
	} else if find := strings.Contains(text, "[["); find {
		if find := strings.Contains(text, "]]"); find {
			//获取主Wiki名字
			Config := utils.ReadConfig()
			MainWikiName := Config.Wiki.([]interface{})[0].(map[interface{}]interface{})["WikiName"].(string)

			trimStr := strings.Trim(text, "[")
			Text := strings.Trim(trimStr, "]")
			return true, Text, MainWikiName
		}
	} else if find := strings.Contains(text, "/"); find {
		SettingsMessageProcessing(json)
		return false, "", "/"
	}

	return false, "", ""
}

//发送群组消息
func sendGroupWikiInfo(WikiName string, GroupID int, QueryText string, quoteID int) {
	WikiInfo, err := Plugin.GetWikiInfo(WikiName, QueryText)
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		SendGroupMessage(GroupID, Error(WikiLink), true, quoteID)
		return
	}
	SendGroupMessage(GroupID, WikiInfo, true, quoteID)
}

//发送好友消息
func sendFriendWikiInfo(WikiName string, UserID int, QueryText string) {
	WikiInfo, err := Plugin.GetWikiInfo(WikiName, QueryText)
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		SendFriendMessage(UserID, Error(WikiLink), false, 0)
		return
	}
	SendFriendMessage(UserID, WikiInfo, false, 0)
}

//发送临时会话消息
func sendTempdWikiInfo(WikiName string, UserID int, GroupID int, QueryText string) {
	WikiInfo, err := Plugin.GetWikiInfo(WikiName, QueryText)
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		SendTempMessage(UserID, GroupID, Error(WikiLink), false, 0)
		return
	}
	SendTempMessage(UserID, GroupID, WikiInfo, false, 0)
}

//戳一戳消息处理
func NudgeEventMessageProcessing(json Struct.QQWebHook_root) {
	HelpText := Language.Message().HelpText
	switch json.Subject.Kind {
	case "Group":
		if json.FromId != utils.ReadConfig().QQBot.BotQQNumber && json.Target == utils.ReadConfig().QQBot.BotQQNumber {
			go SendNudge(json.FromId, json.Subject.Id, "Group")
			go SendGroupAtMessage(json.Subject.Id, HelpText, json.FromId)
		}
	case "Friend":
		go SendFriendMessage(json.FromId, HelpText, false, 0)
	}
}

//消息处理
func MessageProcessing(json Struct.QQWebHook_root) {
	switch json.Type {
	case "GroupMessage":
		if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
			text := json.MessageChain[1].(map[string]interface{})["text"]
			find, QueryText, Command := CommandExtraction(json, text.(string))
			if find {
				GroupID := json.Sender.Group.Id
				quoteID := int(math.Floor(json.MessageChain[0].(map[string]interface{})["id"].(float64)))
				UserID := json.Sender.Id
				go SendNudge(UserID, GroupID, "Group")
				go sendGroupWikiInfo(Command, GroupID, QueryText, quoteID)
			}
		}
	case "FriendMessage":
		if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
			text := json.MessageChain[1].(map[string]interface{})["text"]
			find, QueryText, Command := CommandExtraction(json, text.(string))
			if find {
				UserID := json.Sender.Id
				go sendFriendWikiInfo(Command, UserID, QueryText)
			}
		}
	case "TempMessage":
		if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
			text := json.MessageChain[1].(map[string]interface{})["text"]
			find, QueryText, Command := CommandExtraction(json, text.(string))
			if find {
				UserID := json.Sender.Id
				GroupID := json.Sender.Group.Id
				go sendTempdWikiInfo(Command, UserID, GroupID, QueryText)
			}
		}
	case "NudgeEvent":
		NudgeEventMessageProcessing(json)
	}
}

//设置消息返回
func SettingsMessageProcessing(json Struct.QQWebHook_root) {
	text := json.MessageChain[1].(map[string]interface{})["text"]
	countSplit := strings.Split(text.(string), "/")
	Text := countSplit[1]
	Message, Bool := Settings.Settings(json, Text)
	if Bool {
		switch json.Type {
		case "GroupMessage":
			GroupID := json.Sender.Group.Id
			quoteID := int(math.Floor(json.MessageChain[0].(map[string]interface{})["id"].(float64)))
			go SendGroupMessage(GroupID, Message, true, quoteID)
		case "FriendMessage":
			UserID := json.Sender.Id
			go SendFriendMessage(UserID, Message, false, 0)
		case "TempMessage":
			UserID := json.Sender.Id
			GroupID := json.Sender.Group.Id
			go SendTempMessage(UserID, GroupID, Message, false, 0)
		}
	}
}
