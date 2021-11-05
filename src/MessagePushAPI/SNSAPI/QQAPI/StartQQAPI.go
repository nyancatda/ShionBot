package QQAPI

import (
	"fmt"
	"time"

	"xyz.nyan/MediaWiki-Bot/src/utils/Language"
)

//定时请求mirai-api-http Session
func CycleGetKey() {
	for {
		timer := time.NewTimer(1 * time.Second)
		<-timer.C
		time.Sleep(299 * time.Second)
		_, resp, err := CreateSessionKey()
		if err != nil {
			fmt.Println(Language.Message("", "").UnableApplySession)
			fmt.Println(err)
		} else if resp.Status != "200 OK" {
			fmt.Println(Language.Message("", "").UnableApplySession)
		}
	}
}

func StartQQAPI() error {
	//缓存mirai-api-http Session并启动定时获取进程
	_, resp, err := CreateSessionKey()
	if err != nil {
		fmt.Println(Language.Message("", "").CannotConnectMirai)
		return err
	} else if resp.Status != "200 OK" {
		fmt.Println(Language.Message("", "").CannotConnectMirai)
		return err
	}
	go CycleGetKey()
	return err
}
