package QQInformationProcessing

import (
	"math"
	"strings"
	"xyz.nyan/MediaWiki-Bot/Plugin"
)

type WebHook_root struct {
	Type         string        `json:"type"`
	Sender       SenderJson    `json:"sender"`
	MessageChain []interface{} `json:"messageChain"`
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

//消息处理,这里判断是哪类消息
func MessageProcessing(json WebHook_root) {
	switch json.Type {
	case "GroupMessage":
		GroupMessageProcessing(json)
	case "FriendMessage":
		FriendMessageProcessing(json)
	case "TempMessage":
		TempMessageProcessing(json)
	}
}

//命令处理，判断命令是否匹配，匹配则输出命令参数
func CommandExtraction(text string) (bool,string) {
	if find := strings.Contains(text, "mw:"); find {
		countSplit := strings.SplitN(text, ":", 2)
		QueryText := countSplit[1]
		return true,QueryText
	}
	return false,""
}

func sendGroupWikiInfo(GroupID int, QueryText string, quoteID int) {
	WikiInfo := Plugin.GetWikiInfo(QueryText)
	SendGroupMessage(GroupID, WikiInfo, true, quoteID)
}

//群消息处理
func GroupMessageProcessing(json WebHook_root) {
	//只处理文字消息
	if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
		text := json.MessageChain[1].(map[string]interface{})["text"]
		find,QueryText := CommandExtraction(text.(string))
		if find {
			GroupID := json.Sender.Group.Id
			quoteID := int(math.Floor(json.MessageChain[0].(map[string]interface{})["id"].(float64)))
			UserID := json.Sender.Id
			go SendNudge(UserID, GroupID, "Group")
			go sendGroupWikiInfo(GroupID, QueryText, quoteID)
		}
	}
}

func sendFriendWikiInfo(UserID int, QueryText string) {
	WikiInfo := Plugin.GetWikiInfo(QueryText)
	SendFriendMessage(UserID,WikiInfo,false,0)
}

//好友消息处理
func FriendMessageProcessing(json WebHook_root) {
	if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
		text := json.MessageChain[1].(map[string]interface{})["text"]
		find,QueryText := CommandExtraction(text.(string))
		if find {
			UserID := json.Sender.Id
			go sendFriendWikiInfo(UserID, QueryText)
		}
	}
}

func sendTempdWikiInfo(UserID int, GroupID int,QueryText string) {
	WikiInfo := Plugin.GetWikiInfo(QueryText)
	SendTempMessage(UserID,GroupID,WikiInfo,false,0)
}

//临时会话消息处理
func TempMessageProcessing(json WebHook_root) {
	if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
		text := json.MessageChain[1].(map[string]interface{})["text"]
		find,QueryText := CommandExtraction(text.(string))
		if find {
			UserID := json.Sender.Id
			GroupID := json.Sender.Group.Id
			go sendTempdWikiInfo(UserID, GroupID,QueryText)
		}
	}
}