package QQInformationProcessing

import (
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
	Group              GroupJson `json:"group"`
}
type GroupJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func sendWikiInfo(GroupID int, QueryText string) {
	WikiInfo := Plugin.GetWikiInfo(QueryText)
	SendGroupMessage(GroupID, WikiInfo)
}

//群消息处理
func GroupMessageProcessing(json WebHook_root) {
	//只处理文字消息
	if json.MessageChain[1].(map[string]interface{})["type"] == "Plain" {
		text := json.MessageChain[1].(map[string]interface{})["text"]
		if find := strings.Contains(text.(string), "mw:"); find {
			countSplit := strings.SplitN(text.(string), ":", 2)
			QueryText := countSplit[1]
			GroupID := json.Sender.Group.Id
			go sendWikiInfo(GroupID, QueryText)
		}
	}
}