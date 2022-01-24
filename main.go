package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nyancatda/ShionBot/src/HttpAPI"
	"github.com/nyancatda/ShionBot/src/InformationProcessing"
	"github.com/nyancatda/ShionBot/src/MessagePushAPI/SNSAPI/QQAPI"
	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/Utils"
	"github.com/nyancatda/ShionBot/src/Utils/Language"
	"github.com/nyancatda/ShionBot/src/Utils/ReleaseFile"
)

func Error() {
	fmt.Printf(Language.DefaultLanguageMessage().MainErrorTips)
	key := make([]byte, 1)
	os.Stdin.Read(key)
	os.Exit(1)
}

func main() {
	//参数解析
	ConfigPath := flag.String("config", "./config.yml", "指定配置文件路径")
	flag.Parse()

	//释放资源文件
	ReleaseFile.ReleaseFile()

	//建立数据储存文件夹
	_, err := os.Stat("./data")
	if err != nil {
		os.MkdirAll("./data", 0777)
	}

	//设置配置文件路径
	Utils.ConfigPath = *ConfigPath
	//加载配置文件
	if err := Utils.LoadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Config := Utils.GetConfig

	//判断是否需要初始化QQ部分
	if Config.SNS.QQ.Switch {
		QQAPI.StartQQAPI()
	}

	//启动WebHook接收
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	Port := Config.Run.WebHookPort
	fmt.Println(Utils.StringVariable(Language.DefaultLanguageMessage().RunOK, []string{Port}))
	WebHookKey := Config.Run.WebHookKey
	r.POST("/"+WebHookKey, func(c *gin.Context) {
		var json Struct.WebHookJson
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		InformationProcessing.InformationProcessing(c, json)
	})

	//启动API
	HttpAPI.HttpAPIStart(r)

	r.Run(":" + Port)
}
