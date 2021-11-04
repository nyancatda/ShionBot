package SNSAPI

import (
	"time"
	"fmt"
)

//日志输出
//Type 消息类型，可选 Friend,Group,Stranger
//target 消息接收者
//text 消息主体
func Log(Type string, target string, text string) {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)

	fmt.Println("[" + tm.Format("2006-01-02 03:04:05") + "] [" + Type + "] " + target + " -> " + text)
}