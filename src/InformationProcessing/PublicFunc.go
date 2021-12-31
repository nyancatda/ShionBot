package InformationProcessing

import (
	"fmt"
	"strings"
	"time"

	"github.com/nyancatda/ShionBot/src/Modular/GetWikiInfo"
	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/utils"
	"github.com/nyancatda/ShionBot/src/utils/Language"
)

//Wiki链接错误返回
func Error(SNSName string, UserID string, WikiLink string) string {
	return utils.StringVariable(Language.Message(SNSName, UserID).WikiLinkError, []string{WikiLink})
}

//消息处理日志
//SNSName 聊天软件名字
//Type 消息来源类型
//target 消息发送者
//text 消息主体
func Log(SNSName string, Type string, target string, text string) {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)

	LogText := "[" + tm.Format("2006-01-02 03:04:05") + "] [" + SNSName + "] [" + Type + "] " + target + " -> " + text
	fmt.Println(LogText)

	utils.LogWrite(LogText)
}

//命令处理，判断命令是否匹配，匹配则输出命令和命令参数
func CommandExtraction(SNSName string, json Struct.WebHookJson, text string) (bool, string, string) {
	if find := strings.Contains(text, ":"); find {
		if GetWikiInfo.WikiNameExist(text, SNSName, json) {
			countSplit := strings.SplitN(text, ":", 2)
			Command := countSplit[0]
			Text := countSplit[1]
			return true, Text, Command
		}
	} else if find := strings.Contains(text, "[["); find {
		if find := strings.Contains(text, "]]"); find {
			//获取主Wiki名字
			MainWikiName := GetWikiInfo.GeiMainWikiName(SNSName, json)

			trimStr := strings.Trim(text, "[")
			Text := strings.Trim(trimStr, "]")
			return true, Text, MainWikiName
		}
	} else if find := strings.Contains(text, "/"); find {
		switch SNSName {
		case "QQ":
			QQSettingsMessageProcessing(json)
		case "Telegram":
			TelegramSettingsMessageProcessing(json)
		case "Line":
			LineSettingsMessageProcessing(json)
		case "KaiHeiLa":
			KaiHeiLaSettingsMessageProcessing(json)
		}
		return false, "", "/"
	}

	return false, "", ""
}
