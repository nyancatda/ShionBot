package MessageProcessing

import (
	"xyz.nyan/MediaWiki-Bot/QQInformationProcessing"
)

//发送群消息
//SNSName 聊天软件名字
//target 群号
//text 消息文本
//quote 是否需要回复
//quoteID 回复的消息ID(不需要时为0即可)
func SendGroupMessage(SNSName string, target int, text string, quote bool, quoteID int) {
	if SNSName == "QQ" {
		QQInformationProcessing.SendGroupMessage(target, text, quote, quoteID)
	}
}
