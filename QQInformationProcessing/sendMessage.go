package QQInformationProcessing

import (
	"fmt"
	"strconv"
	"time"

	"xyz.nyan/MediaWiki-Bot/utils"
)

//日志输出
//Type 消息类型，可选 Friend,Group,Stranger
//target 消息来源
//text 消息主体
func log(Type string, target string, text string) {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)

	fmt.Println("[" + tm.Format("2006-01-02 03:04:05") + "] [" + Type + "] " + target + " -> " + text)
}

//发送群消息
//target 群号
//text 消息文本
//quote 是否需要回复
//quoteID 回复的消息ID(不需要时为0即可)
func SendGroupMessage(target int, text string, quote bool, quoteID int) {
	Config := utils.ReadConfig()
	sessionKey := GetSessionKey()
	var requestBody string
	//判断是否需要引用回复
	if quote {
		requestBody = fmt.Sprintf(`{
			"sessionKey": "%s",
			"target": %d,
			"quote": %d,
			"messageChain": [
			  {
				"type": "Plain",
				"text": "%s"
			  }
			]
		}`, sessionKey, target, quoteID, text)
	} else {
		requestBody = fmt.Sprintf(`{
			"sessionKey": "%s",
			"target": %d,
			"messageChain": [
			  {
				"type": "Plain",
				"text": "%s"
			  }
			]
		}`, sessionKey, target, text)
	}

	url := Config.QQBot.APILink + "/sendGroupMessage"
	utils.PostRequestJosn(url, requestBody)
	log("Group",strconv.Itoa(target), text)
}

//发送头像戳一戳
//target 目标QQ号
//subject 消息接受主体，为群号或QQ号
//kind 上下文类型,可选值 Friend,Group,Stranger
func SendNudge(target int, subject int, kind string) {
	Config := utils.ReadConfig()
	sessionKey := GetSessionKey()
	requestBody := fmt.Sprintf(`{
		"sessionKey":"%s",
		"target":%d,
		"subject":%d,
		"kind":"%s"
	}`, sessionKey, target, subject, kind)

	url := Config.QQBot.APILink + "/sendNudge"
	utils.PostRequestJosn(url, requestBody)
	log(kind,strconv.Itoa(subject), "戳一戳"+strconv.Itoa(target))
}
