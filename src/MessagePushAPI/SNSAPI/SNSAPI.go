/*
 * @Author: NyanCatda
 * @Date: 2021-11-05 13:51:15
 * @LastEditTime: 2022-01-24 20:03:33
 * @LastEditors: NyanCatda
 * @Description: 聊天软件API封装
 * @FilePath: \ShionBot\src\MessagePushAPI\SNSAPI\SNSAPI.go
 */
package SNSAPI

import (
	"fmt"
	"time"

	"github.com/nyancatda/ShionBot/src/Utils"
)

/**
 * @description: 日志输出
 * @param {string} SNSName 聊天软件名字
 * @param {string} Type 消息类型，可选 Friend,Group,Stranger,Channel
 * @param {string} target 消息接收者
 * @param {string} text 消息主体
 * @return {*}
 */
func Log(SNSName string, Type string, target string, text string) {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)

	LogText := "[" + tm.Format("2006-01-02 03:04:05") + "] [" + SNSName + "] [" + Type + "] " + target + " <- " + text
	fmt.Println(LogText)

	Utils.LogWrite(LogText)
}
