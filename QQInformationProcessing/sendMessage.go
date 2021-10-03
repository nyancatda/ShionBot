package QQInformationProcessing

import (
	"fmt"
	"strconv"
	"time"
	"xyz.nyan/MediaWiki-Bot/utils"
)

func log(target string, text string) {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)

	fmt.Println("[" + tm.Format("2006-01-02 03:04:05") + "] " + target + " -> " + text)
}

//发送群消息
//target 群号
//text 消息文本
func SendGroupMessage(target int, text string) {
	Config := utils.ReadConfig()
	sessionKey := GetSessionKey()
	requestBody := fmt.Sprintf(`{
        "sessionKey": "%s",
        "target": %d,
        "messageChain": [
          {
            "type": "Plain",
            "text": "%s"
          }
        ]
    }`, sessionKey, target, text)
	url := Config.QQBot.APILink + "/sendGroupMessage"
	utils.PostRequestJosn(url, requestBody)
	log(strconv.Itoa(target),text)
}
