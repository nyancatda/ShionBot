package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"xyz.nyan/MediaWiki-Bot/src/InformationProcessing"
	"xyz.nyan/MediaWiki-Bot/src/MessageProcessingAPI/SNSAPI/QQAPI"
	"xyz.nyan/MediaWiki-Bot/src/Struct"
	"xyz.nyan/MediaWiki-Bot/src/utils"
	"xyz.nyan/MediaWiki-Bot/src/utils/Language"
	"xyz.nyan/MediaWiki-Bot/src/utils/ReleaseFile"
)

func Error() {
	fmt.Printf(Language.Message("").MainErrorTips)
	key := make([]byte, 1)
	os.Stdin.Read(key)
	os.Exit(1)
}

//定时请求mirai-api-http Session
func CycleGetKey() {
	for {
		timer := time.NewTimer(1 * time.Second)
		<-timer.C
		time.Sleep(299 * time.Second)
		_, resp, err := QQAPI.CreateSessionKey()
		if err != nil {
			fmt.Println(Language.Message("").UnableApplySession)
			fmt.Println(err)
		} else if resp.Status != "200 OK" {
			fmt.Println(Language.Message("").UnableApplySession)
		}
	}
}

func main() {
	//释放资源文件
	ReleaseFile.ReleaseFile()

	//读取配置文件
	Config := utils.ReadConfig()

	//缓存mirai-api-http Session并启动定时获取进程
	_, resp, err := QQAPI.CreateSessionKey()
	if err != nil {
		fmt.Println(Language.Message("").CannotConnectMirai)
		Error()
	} else if resp.Status != "200 OK" {
		fmt.Println(Language.Message("").CannotConnectMirai)
		Error()
	}
	go CycleGetKey()

	//启动WebHook接收
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	Port := Config.Run.WebHookPort
	fmt.Println(Language.StringVariable(1, Language.Message("").RunOK, Port, ""))
	r.POST("/", func(c *gin.Context) {
		var json Struct.WebHookJson
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		InformationProcessing.InformationProcessing(json)
	})
	r.Run(":" + Port)
}
