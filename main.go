package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"xyz.nyan/MediaWiki-Bot/QQInformationProcessing"
	"xyz.nyan/MediaWiki-Bot/utils"
)

func main() {
	//判断配置文件是否正常
	if utils.CheckConfigFile() {
		fmt.Println("配置文件异常")
		fmt.Printf("Press any key to exit...")
		key := make([]byte, 1)
		os.Stdin.Read(key)
		os.Exit(1)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	fmt.Println("机器人已运行在8000端口")

	r.POST("/", func(c *gin.Context) {
		var json QQInformationProcessing.WebHook_root
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		QQInformationProcessing.GroupMessageProcessing(json)
	})

	r.Run(":8000")
}
