package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"xyz.nyan/MediaWiki-Bot/QQInformationProcessing"
	"xyz.nyan/MediaWiki-Bot/Struct"
	"xyz.nyan/MediaWiki-Bot/utils"
	"xyz.nyan/MediaWiki-Bot/utils/Language"
)

func Error() {
	fmt.Printf(Language.Message("").MainErrorTips)
	key := make([]byte, 1)
	os.Stdin.Read(key)
	os.Exit(1)
}

func CycleGetKey() {
	for {
		timer := time.NewTimer(1 * time.Second)
		<-timer.C
		time.Sleep(299 * time.Second)
		_, resp, err := QQInformationProcessing.CreateSessionKey()
		if err != nil {
			fmt.Println(Language.Message("").UnableApplySession)
			fmt.Println(err)
		} else if resp.Status != "200 OK" {
			fmt.Println(Language.Message("").UnableApplySession)
		}
	}
}

func main() {
	//释放语言文件
	Language.ReleaseFile()
	//判断配置文件是否正常
	if utils.CheckConfigFile() {
		fmt.Println(Language.Message("").ConfigFileException)
		Error()
	}
	Config := utils.ReadConfig()
	Port := Config.Run.WebHookPort

	_, resp, err := QQInformationProcessing.CreateSessionKey()
	if err != nil {
		fmt.Println(Language.Message("").CannotConnectMirai)
		Error()
	} else if resp.Status != "200 OK" {
		fmt.Println(Language.Message("").CannotConnectMirai)
		Error()
	}

	go CycleGetKey()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	fmt.Println(Language.StringVariable(1, Language.Message("").RunOK, Port, ""))

	r.POST("/", func(c *gin.Context) {
		var json Struct.QQWebHook_root
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		QQInformationProcessing.MessageProcessing(json)
	})

	r.Run(":" + Port)
}
