/*
 * @Author: NyanCatda
 * @Date: 2021-11-04 22:09:03
 * @LastEditTime: 2022-01-27 18:05:47
 * @LastEditors: NyanCatda
 * @Description: Telegram消息处理
 * @FilePath: \ShionBot\src\MessageProcessing\TelegramMessageProcessing.go
 */
package MessageProcessing

import (
	"strconv"

	"github.com/nyancatda/ShionBot/src/MessagePushAPI"
	"github.com/nyancatda/ShionBot/src/Modular/Command"
	"github.com/nyancatda/ShionBot/src/Modular/GetWikiInfo"
	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/Utils/ReadConfig"
)

var sns_name_telegram string = "Telegram"

func TelegramMessageProcessing(json Struct.WebHookJson) {
	text := json.Message.Text
	//判断命令是否匹配
	find, Command, CommandData := CommandExtraction(sns_name_telegram, json, text)
	if find {
		if Command == "/" {
			TelegramSettingsMessageProcessing(CommandData, json)
			return
		}

		ChatType := json.Message.Chat.Type
		UserID := strconv.Itoa(json.Message.From.Id)
		ChatID := strconv.Itoa(json.Message.Chat.Id)
		WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_telegram, json, UserID, Command, CommandData, "")
		if err != nil {
			WikiLink := ReadConfig.GetWikiLink(sns_name_telegram, json, Command)
			MessagePushAPI.SendMessage(sns_name_telegram, "Default", UserID, ChatID, Error(sns_name_telegram, UserID, WikiLink), false, "", "", 0)
			return
		}
		switch ChatType {
		case "private":
			MessagePushAPI.SendMessage(sns_name_telegram, "Default", UserID, ChatID, WikiInfo, false, "", "", 0)
		case "group":
			MassageID := strconv.Itoa(json.Message.Message_id)
			MessagePushAPI.SendMessage(sns_name_telegram, "Group", UserID, ChatID, WikiInfo, true, MassageID, "", 0)
		case "supergroup":
			MassageID := strconv.Itoa(json.Message.Message_id)
			MessagePushAPI.SendMessage(sns_name_telegram, "Group", UserID, ChatID, WikiInfo, true, MassageID, "", 0)
		}
	}
}

//设置消息处理
func TelegramSettingsMessageProcessing(Text string, json Struct.WebHookJson) {
	Message, Bool := Command.Command(sns_name_telegram, json, Text)
	if Bool {
		ChatID := strconv.Itoa(json.Message.Chat.Id)
		UserID := strconv.Itoa(json.Message.From.Id)
		switch json.Message.Chat.Type {
		case "private":
			MessagePushAPI.SendMessage(sns_name_telegram, "Default", UserID, ChatID, Message, false, "", "", 0)
		case "group":
			MassageID := strconv.Itoa(json.Message.Message_id)
			MessagePushAPI.SendMessage(sns_name_telegram, "Group", UserID, ChatID, Message, true, MassageID, "", 0)
		case "supergroup":
			MassageID := strconv.Itoa(json.Message.Message_id)
			MessagePushAPI.SendMessage(sns_name_telegram, "Group", UserID, ChatID, Message, true, MassageID, "", 0)
		}
	}
}
