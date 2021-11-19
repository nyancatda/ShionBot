package SNSAPI

import (
	"fmt"
	"time"
	"xyz.nyan/ShionBot/src/utils"
)

//日志输出
//SNSName 聊天软件名字
//Type 消息类型，可选 Friend,Group,Stranger,Channel
//target 消息接收者
//text 消息主体
func Log(SNSName string, Type string, target string, text string) {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)

	LogText := "[" + tm.Format("2006-01-02 03:04:05") + "] [" + SNSName + "] [" + Type + "] " + target + " <- " + text
	fmt.Println(LogText)

	utils.LogWrite(LogText)
}
