package QQInformationProcessing

import (
	"fmt"

	"xyz.nyan/MediaWiki-Bot/utils"
)

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
}
