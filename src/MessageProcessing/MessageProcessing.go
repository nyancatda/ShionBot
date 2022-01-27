package MessageProcessing

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nyancatda/ShionBot/src/Modular/GetWikiInfo"
	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/Utils"
	"github.com/nyancatda/ShionBot/src/Utils/Language"
)

/**
 * @description: 消息处理分配
 * @param {*gin.Context} c
 * @param {Struct.WebHookJson} json
 * @return {*}
 */
func MessageProcessing(c *gin.Context, json Struct.WebHookJson) {
	if json.Type != "" {
		go QQMessageProcessing(json)
		if len(json.MessageChain) >= 1 {
			Log("QQ", json.Type, strconv.Itoa(json.Sender.Id), json.MessageChain[1].Text)
		}
		return
	}

	if json.Update_id != 0 {
		go TelegramMessageProcessing(json)
		Log("Telegram", json.Message.Chat.Type, strconv.Itoa(json.Message.From.Id), json.Message.Text)
	}

	if json.Destination != "" {
		go LineMessageProcessing(json)
		Log("Line", json.Events[0].Source.Type, json.Events[0].Source.UserId, json.Events[0].Message.Text)
	}

	if json.D.Challenge != "" {
		KaiHeiLaWebHookVerifyProcessing(c, json)
	}
	if json.D.Content != "" {
		go KaiHeiLaMessageProcessing(json)
		Log("KaiHeiLa", json.D.Channel_type, json.D.Author_id, json.D.Content)
	}
}

/**
 * @description: Wiki链接错误返回
 * @param {string} SNSName
 * @param {string} UserID
 * @param {string} WikiLink
 * @return {*}
 */
func Error(SNSName string, UserID string, WikiLink string) string {
	return Utils.StringVariable(Language.Message(SNSName, UserID).WikiLinkError, []string{WikiLink})
}

/**
 * @description: 消息接收日志
 * @param {string} SNSName 聊天软件名字
 * @param {string} Type 消息来源类型
 * @param {string} target 消息发送者
 * @param {string} text 消息主体
 * @return {*}
 */
func Log(SNSName string, Type string, target string, text string) {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)

	LogText := "[" + tm.Format("2006-01-02 03:04:05") + "] [" + SNSName + "] [" + Type + "] " + target + " -> " + text
	fmt.Println(LogText)

	Utils.LogWrite(LogText)
}

/**
 * @description: 命令处理，判断命令是否匹配，匹配则输出命令和命令参数
 * @param {string} SNSName 聊天软件名字
 * @param {Struct.WebHookJson} json 消息主体
 * @param {string} text 消息内容
 * @return {bool} 是否匹配到命令
 * @return {string} 命令头
 * @return {string} 命令内容
 */
func CommandExtraction(SNSName string, json Struct.WebHookJson, text string) (bool, string, string) {
	if find := strings.Contains(text, ":"); find {
		if GetWikiInfo.WikiNameExist(text, SNSName, json) {
			countSplit := strings.SplitN(text, ":", 2)
			Command := countSplit[0]
			Text := countSplit[1]
			return true, Command, Text
		}
	} else if find := strings.Contains(text, "[["); find {
		if find := strings.Contains(text, "]]"); find {
			//获取主Wiki名字
			MainWikiName := GetWikiInfo.GeiMainWikiName(SNSName, json)

			trimStr := strings.Trim(text, "[")
			Text := strings.Trim(trimStr, "]")
			return true, MainWikiName, Text
		}
	} else if find := strings.Contains(text, "/"); find {
		countSplit := strings.SplitN(text, "/", 2)
		Text := countSplit[1]
		return true, "/", Text
	}

	return false, "", ""
}
