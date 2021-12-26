package QQAPI

import (
	"fmt"
	"time"

	"github.com/nyancatda/ShionBot/src/utils/Language"
)

//定时请求mirai-api-http Session
func CycleGetKey() {
	for {
		timer := time.NewTimer(1 * time.Second)
		<-timer.C
		time.Sleep(299 * time.Second)
		_, resp, err := CreateSessionKey()
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
	_, resp, err := CreateSessionKey()
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
