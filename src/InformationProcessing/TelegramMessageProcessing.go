/*
 * @Author: NyanCatda
 * @Date: 2021-11-04 22:09:03
 * @LastEditTime: 2022-01-24 18:10:08
 * @LastEditors: NyanCatda
 * @Description: Telegram消息处理
 * @FilePath: \ShionBot\src\InformationProcessing\TelegramMessageProcessing.go
 */
package InformationProcessing

import (
	"strconv"
	"strings"

	"github.com/nyancatda/ShionBot/src/MessagePushAPI"
	"github.com/nyancatda/ShionBot/src/Modular/Command"
	"github.com/nyancatda/ShionBot/src/Modular/GetWikiInfo"
	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/Utils"
)

var sns_name_telegram string = "Telegram"

func TelegramMessageProcessing(json Struct.WebHookJson) {
	text := json.Message.Text
	find, QueryText, Command := CommandExtraction(sns_name_telegram, json, text)
	if find {
		ChatType := json.Message.Chat.Type
		UserID := strconv.Itoa(json.Message.From.Id)
		Log(sns_name_telegram, ChatType, UserID, text)
		ChatID := strconv.Itoa(json.Message.Chat.Id)
		WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_telegram, json, UserID, Command, QueryText, "")
		if err != nil {
			WikiLink := Utils.GetWikiLink(sns_name_telegram, json, Command)
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

//设置消息返回
func TelegramSettingsMessageProcessing(json Struct.WebHookJson) {
	text := json.Message.Text
	countSplit := strings.SplitN(text, "/", 2)
	Text := countSplit[1]
	Message, Bool := Command.Command(sns_name_telegram, json, Text)
	if Bool {
		ChatID := strconv.Itoa(json.Message.Chat.Id)
		ChatType := json.Message.Chat.Type
		UserID := strconv.Itoa(json.Message.From.Id)
		Log(sns_name_telegram, ChatType, UserID, text)
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
