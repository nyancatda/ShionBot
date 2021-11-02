package QQInformationProcessing

import (
	"math"
	"strings"

	"xyz.nyan/MediaWiki-Bot/MediaWikiAPI"
	"xyz.nyan/MediaWiki-Bot/Plugin"
	"xyz.nyan/MediaWiki-Bot/utils"
	"xyz.nyan/MediaWiki-Bot/utils/Language"
)

type WebHook_root struct {
	Type         string        `json:"type"`
	Sender       SenderJson    `json:"sender"`
	FromId       int           `json:"fromId"`
	Target       int           `json:"target"`
	MessageChain []interface{} `json:"messageChain"`
	Subject      SubjectJson   `json:"subject"`
}
type SenderJson struct {
	Id                 int       `json:"id"`
	MemberName         string    `json:"memberName"`
	SpecialTitle       string    `json:"specialTitle"`
	Permission         string    `json:"permission"`
	JoinTimestamp      int       `json:"joinTimestamp"`
	LastSpeakTimestamp int       `json:"lastSpeakTimestamp"`
	MuteTimeRemaining  int       `json:"muteTimeRemaining"`
	Nickname           string    `json:"nickname"`
	Remark             string    `json:"remark"`
	Group              GroupJson `json:"group"`
}
type GroupJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type SubjectJson struct {
	Id   int    `json:"id"`
	Kind string `json:"kind"`
}

func Error(WikiLink string) string {
	return Language.StringVariable(1, Language.Message().WikiLinkError, WikiLink, "")
}

//消息处理,这里判断是哪类消息
func MessageProcessing(json WebHook_root) {
	switch json.Type {
	case "GroupMessage":
		GroupMessageProcessing(json)
	case "FriendMessage":
		FriendMessageProcessing(json)
	case "TempMessage":
		TempMessageProcessing(json)
	case "NudgeEvent":
		NudgeEventMessageProcessing(json)
	}
}

//命令处理，判断命令是否匹配，匹配则输出命令和命令参数
func CommandExtraction(text string) (bool, string, string) {
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
	}

	return false, "", ""
}

func sendGroupWikiInfo(WikiName string, GroupID int, QueryText string, quoteID int) {
	WikiInfo, err := Plugin.GetWikiInfo(WikiName, QueryText)
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		SendGroupMessage(GroupID, Error(WikiLink), true, quoteID)
		return
	}
	SendGroupMessage(GroupID, WikiInfo, true, quoteID)
}

//群消息处理
func GroupMessageProcessing(json WebHook_root) {
	//只处理文字消息
	if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
		text := json.MessageChain[1].(map[string]interface{})["text"]
		find, QueryText, Command := CommandExtraction(text.(string))
		if find {
			GroupID := json.Sender.Group.Id
			quoteID := int(math.Floor(json.MessageChain[0].(map[string]interface{})["id"].(float64)))
			UserID := json.Sender.Id
			go SendNudge(UserID, GroupID, "Group")
			go sendGroupWikiInfo(Command, GroupID, QueryText, quoteID)
		}
	}
}

func sendFriendWikiInfo(WikiName string, UserID int, QueryText string) {
	WikiInfo, err := Plugin.GetWikiInfo(WikiName, QueryText)
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		SendFriendMessage(UserID, Error(WikiLink), false, 0)
		return
	}
	SendFriendMessage(UserID, WikiInfo, false, 0)
}

//好友消息处理
func FriendMessageProcessing(json WebHook_root) {
	if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
		text := json.MessageChain[1].(map[string]interface{})["text"]
		find, QueryText, Command := CommandExtraction(text.(string))
		if find {
			UserID := json.Sender.Id
			go sendFriendWikiInfo(Command, UserID, QueryText)
		}
	}
}

func sendTempdWikiInfo(WikiName string, UserID int, GroupID int, QueryText string) {
	WikiInfo, err := Plugin.GetWikiInfo(WikiName, QueryText)
	if err != nil {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		SendTempMessage(UserID, GroupID, Error(WikiLink), false, 0)
		return
	}
	SendTempMessage(UserID, GroupID, WikiInfo, false, 0)
}

//临时会话消息处理
func TempMessageProcessing(json WebHook_root) {
	if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
		text := json.MessageChain[1].(map[string]interface{})["text"]
		find, QueryText, Command := CommandExtraction(text.(string))
		if find {
			UserID := json.Sender.Id
			GroupID := json.Sender.Group.Id
			go sendTempdWikiInfo(Command, UserID, GroupID, QueryText)
		}
	}
}

func NudgeEventMessageProcessing(json WebHook_root) {
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
