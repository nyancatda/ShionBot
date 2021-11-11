package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"xyz.nyan/MediaWiki-Bot/src/HttpAPI"
	"xyz.nyan/MediaWiki-Bot/src/InformationProcessing"
	"xyz.nyan/MediaWiki-Bot/src/MessagePushAPI/SNSAPI/QQAPI"
	"xyz.nyan/MediaWiki-Bot/src/Struct"
	"xyz.nyan/MediaWiki-Bot/src/utils"
	"xyz.nyan/MediaWiki-Bot/src/utils/Language"
	"xyz.nyan/MediaWiki-Bot/src/utils/ReleaseFile"
)

func Error() {
	fmt.Printf(Language.DefaultLanguageMessage().MainErrorTips)
	key := make([]byte, 1)
	os.Stdin.Read(key)
	os.Exit(1)
}

func main() {
	//释放资源文件
	ReleaseFile.ReleaseFile()

	//建立数据储存文件夹
	_, err := os.Stat("./data")
	if err != nil {
		os.MkdirAll("./data", 0777)
	}

	//读取配置文件
	Config := utils.ReadConfig()

	//判断是否需要初始化QQ部分
	if Config.SNS.QQ.Switch {
		QQAPI.StartQQAPI()
	}

	//启动WebHook接收
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	Port := Config.Run.WebHookPort
	fmt.Println(utils.StringVariable(Language.DefaultLanguageMessage().RunOK, []string{Port}))
	WebHookKey := Config.Run.WebHookKey
	r.POST("/"+WebHookKey, func(c *gin.Context) {
		var json Struct.WebHookJson
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		InformationProcessing.InformationProcessing(c,json)
	})

	//启动API
	HttpAPI.HttpAPIStart(r)

	r.Run(":" + Port)
}
