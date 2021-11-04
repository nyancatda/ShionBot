package MessageProcessingAPI

import (
	"strconv"
	"xyz.nyan/MediaWiki-Bot/src/MessageProcessingAPI/SNSAPI/QQAPI"
	"xyz.nyan/MediaWiki-Bot/src/MessageProcessingAPI/SNSAPI/TelegramAPI"
)

//发送好友消息
//SNSName 聊天软件
//target 好友ID
//text 消息文本
//quote 是否需要回复
//quoteID 回复的消息ID(不需要时为0即可)
func SendFriendMessage(SNSName string, target int, text string, quote bool, quoteID int) {
	switch SNSName {
	case "QQ":
		QQAPI.SendFriendMessage(target, text, quote, quoteID)
	case "Telegram":
		TelegramAPI.SendMessage("Friend", target, text, true, false, quoteID, quote)
	}
}

//发送群消息
//SNSName 聊天软件
//target 群号
//text 消息文本
//quote 是否需要回复
//quoteID 回复的消息ID(不需要时为0即可)
func SendGroupMessage(SNSName string, target int, text string, quote bool, quoteID int) {
	switch SNSName {
	case "QQ":
		QQAPI.SendGroupMessage(target, text, quote, quoteID)
	case "Telegram":
		TelegramAPI.SendMessage("Group", target, text, true, false, quoteID, quote)
	}
}

//发送带@的群消息
//SNSName 聊天软件
//target 群号
//text 消息文本
//AtID 需要@的人的ID
func SendGroupAtMessage(SNSName string, target int, text string, AtID string) {
	switch SNSName {
	case "QQ":
		AtID, _ := strconv.Atoi(AtID)
		QQAPI.SendGroupAtMessage(target, text, AtID)
	case "Telegram":
		text = "@" + AtID + " " + text
		TelegramAPI.SendMessage("Group", target, text, true, false, 0, false)
	}
}

//发送陌生人会话
//SNSName 聊天软件
//target 陌生人对象ID
//group 临时会话群号
//text 消息文本
//quote 是否需要回复
//quoteID 回复的消息ID(不需要时为0即可)
func SendTempMessage(SNSName string, target int, group int, text string, quote bool, quoteID int) {
	switch SNSName {
	case "QQ":
		QQAPI.SendTempMessage(target, group, text, quote, quoteID)
	case "Telegram":
		TelegramAPI.SendMessage("Stranger", target, text, true, false, quoteID, quote)
	}
}

//(QQ)发送头像戳一戳
//target 目标QQ号
//subject 消息接受主体，为群号或QQ号
//kind 上下文类型,可选值 Friend,Group,Stranger
func SendNudge(target int, subject int, kind string) {
	QQAPI.SendNudge(target, subject, kind)
}
