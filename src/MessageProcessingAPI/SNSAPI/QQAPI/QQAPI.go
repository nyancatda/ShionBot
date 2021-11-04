package QQAPI

import (
	"encoding/json"
	"fmt"
	"strconv"

	"xyz.nyan/MediaWiki-Bot/src/utils"
	"xyz.nyan/MediaWiki-Bot/src/utils/Language"
	"xyz.nyan/MediaWiki-Bot/src/MessageProcessingAPI/SNSAPI"
)

type returnJson struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	MessageId string `json:"messageId"`
}

func sendError(body []byte, err error, url string, requestBody string) {
	if err != nil {
		fmt.Println(Language.Message("").CannotConnectMirai)
	} else {
		var config returnJson
		json.Unmarshal([]byte(body), &config)
		if config.Code != 0 {
			SessionKey, resp, err := CreateSessionKey()
			if err != nil {
				fmt.Println(Language.Message("").CannotConnectMirai)
				fmt.Println(err)
			} else if resp.Status != "200 OK" {
				fmt.Println(Language.Message("").CannotConnectMirai)
			} else {
				var result map[string]interface{}
				json.Unmarshal([]byte(requestBody), &result)
				result["sessionKey"] = SessionKey
				Body, _ := json.Marshal(result)
				utils.PostRequestJosn(url, string(Body))
			}
		}
	}
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
	body, _, err := utils.PostRequestJosn(url, requestBody)
	sendError(body, err, url, requestBody)

	SNSAPI.Log("Group", strconv.Itoa(target), text)
}

//发送带@的群消息
//target 群号
//text 消息文本
//AtID 需要@的人的QQ号
func SendGroupAtMessage(target int, text string, AtID int) {
	Config := utils.ReadConfig()
	sessionKey := GetSessionKey()
	//判断是否需要引用回复
	requestBody := fmt.Sprintf(`{
			"sessionKey": "%s",
			"target": %d,
			"messageChain": [
			{
				"type": "At",
				"target": %d
			},
			{
			"type": "Plain",
			"text": "%s"
			}
			]
		}`, sessionKey, target, AtID, text)
	url := Config.QQBot.APILink + "/sendGroupMessage"
	body, _, err := utils.PostRequestJosn(url, requestBody)
	sendError(body, err, url, requestBody)

	SNSAPI.Log("Group", strconv.Itoa(target), text)
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
	body, _, err := utils.PostRequestJosn(url, requestBody)
	sendError(body, err, url, requestBody)

	SNSAPI.Log(kind, strconv.Itoa(subject), Language.Message("").Nudge+strconv.Itoa(target))
}

//发送好友消息
//target 好友QQ号
//text 消息文本
//quote 是否需要回复
//quoteID 回复的消息ID(不需要时为0即可)
func SendFriendMessage(target int, text string, quote bool, quoteID int) {
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

	url := Config.QQBot.APILink + "/sendFriendMessage"
	body, _, err := utils.PostRequestJosn(url, requestBody)
	sendError(body, err, url, requestBody)

	SNSAPI.Log("Friend", strconv.Itoa(target), text)
}

//发送临时会话
//target 临时会话对象QQ号
//group 临时会话群号
//text 消息文本
//quote 是否需要回复
//quoteID 回复的消息ID(不需要时为0即可)
func SendTempMessage(target int, group int, text string, quote bool, quoteID int) {
	Config := utils.ReadConfig()
	sessionKey := GetSessionKey()
	var requestBody string
	//判断是否需要引用回复
	if quote {
		requestBody = fmt.Sprintf(`{
			"sessionKey": "%s",
			"qq": %d,
			"group": %d
			"quote": %d,
			"messageChain": [
			  {
				"type": "Plain",
				"text": "%s"
			  }
			]
		}`, sessionKey, target, group, quoteID, text)
	} else {
		requestBody = fmt.Sprintf(`{
			"sessionKey": "%s",
			"qq": %d,
			"group": %d
			"messageChain": [
			  {
				"type": "Plain",
				"text": "%s"
			  }
			]
		}`, sessionKey, target, group, text)
	}

	url := Config.QQBot.APILink + "/sendTempMessage"
	body, _, err := utils.PostRequestJosn(url, requestBody)
	sendError(body, err, url, requestBody)

	SNSAPI.Log("Temp", strconv.Itoa(target), text)
}
