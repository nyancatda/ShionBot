package MessageProcessingAPI

import (
	"strconv"
	"xyz.nyan/MediaWiki-Bot/src/MessageProcessingAPI/SNSAPI/QQAPI"
	"xyz.nyan/MediaWiki-Bot/src/MessageProcessingAPI/SNSAPI/TelegramAPI"
)

//发送消息
//SNSName 聊天软件
//ChatType 聊天类型(Default,Friend,Group,GroupAt,Temp)
//target 接受的聊天的ID
//text 消息文本
//quote 是否需要回复
//quoteID 回复的消息ID(不需要时为0即可)
//AtID (可选)需要@的人的ID
//group (可选)临时会话群号
func SendMessage(SNSName string, ChatType string, target int, text string, quote bool, quoteID int, AtID string, group int) {
	switch SNSName {
	case "QQ":
		switch ChatType {
		case "Friend":
			QQAPI.SendFriendMessage(target, text, quote, quoteID)
		case "Group":
			QQAPI.SendGroupMessage(target, text, quote, quoteID)
		case "GroupAt":
			AtID, _ := strconv.Atoi(AtID)
			QQAPI.SendGroupAtMessage(target, text, AtID)
		case "Temp":
			QQAPI.SendTempMessage(target, group, text, quote, quoteID)
		}
	case "Telegram":
		switch ChatType {
		case "GroupAt":
			text = "@" + AtID + " " + text
			TelegramAPI.SendMessage("Group", target, text, true, false, 0, false)
		default:
			TelegramAPI.SendMessage("Friend", target, text, true, false, quoteID, quote)
		}
	}
}

//(QQ)发送头像戳一戳
//target 目标QQ号
//subject 消息接受主体，为群号或QQ号
//kind 上下文类型,可选值 Friend,Group,Stranger
func SendNudge(target int, subject int, kind string) {
	QQAPI.SendNudge(target, subject, kind)
}
