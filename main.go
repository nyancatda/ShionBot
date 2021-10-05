package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"xyz.nyan/MediaWiki-Bot/QQInformationProcessing"
	"xyz.nyan/MediaWiki-Bot/utils"
)

func Error() {
	fmt.Printf("Press any key to exit...")
	key := make([]byte, 1)
	os.Stdin.Read(key)
	os.Exit(1)
}

func CycleGetKey() {
	for {
		timer := time.NewTimer(1 * time.Second)
		<-timer.C
		time.Sleep(299 * time.Second)
		_,resp,err := QQInformationProcessing.CreateSessionKey()
		if err != nil {
			fmt.Println("无法向Mirai申请新的Session，请检查设置")
			fmt.Println(err)
		} else if resp.Status != "200 OK" {
			fmt.Println("无法向Mirai申请新的Session，请检查设置")
		}
	}
}

func main() {
	//判断配置文件是否正常
	if utils.CheckConfigFile() {
		fmt.Println("配置文件异常")
		Error()
	}
	Config := utils.ReadConfig()
	Port := Config.Run.WebHookPort

	_,resp,err := QQInformationProcessing.CreateSessionKey()
	if err != nil {
		fmt.Println("无法正确链接至Mirai，请检查设置")
		Error()
	} else if resp.Status != "200 OK" {
		fmt.Println("无法正确链接至Mirai，请检查设置")
		Error()
	}

	go CycleGetKey()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	fmt.Println("机器人WebHook接收已开启，运行于"+Port+"端口")

	r.POST("/", func(c *gin.Context) {
		var json QQInformationProcessing.WebHook_root
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		QQInformationProcessing.MessageProcessing(json)
	})

	r.Run(":"+Port)
}
