/*
 * @Author: NyanCatda
 * @Date: 2021-11-05 18:12:19
 * @LastEditTime: 2022-01-24 21:20:31
 * @LastEditors: NyanCatda
 * @Description: QQ API定时器
 * @FilePath: \ShionBot\src\MessagePushAPI\SNSAPI\QQAPI\StartQQAPI.go
 */
package QQAPI

import (
	"fmt"
	"time"

	"github.com/nyancatda/ShionBot/src/Utils/Language"
)

//定时请求mirai-api-http Session
func CycleGetKey() {
	for {
		timer := time.NewTimer(1 * time.Second)
		<-timer.C
		time.Sleep(299 * time.Second)
		_, _, resp, err := CreateSessionKey()
		if err != nil {
			fmt.Println(Language.DefaultLanguageMessage().UnableApplySession)
			fmt.Println(err)
		} else if resp.Status != "200 OK" {
			fmt.Println(Language.DefaultLanguageMessage().UnableApplySession)
		}
	}
}

func StartQQAPI() error {
	//缓存mirai-api-http Session并启动定时获取进程
	_, _, resp, err := CreateSessionKey()
	if err != nil {
		fmt.Println(Language.DefaultLanguageMessage().CannotConnectMirai)
		return err
	} else if resp.Status != "200 OK" {
		fmt.Println(Language.DefaultLanguageMessage().CannotConnectMirai)
		return err
	}
	go CycleGetKey()
	return err
}
