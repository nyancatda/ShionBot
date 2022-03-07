/*
 * @Author: NyanCatda
 * @Date: 2021-11-05 13:51:15
 * @LastEditTime: 2022-01-24 21:12:06
 * @LastEditors: NyanCatda
 * @Description: 消息发送封装
 * @FilePath: \ShionBot\src\MessagePushAPI\SendMessage.go
 */
package MessagePushAPI

import (
	"net/http"
	"strconv"

	"github.com/nyancatda/ShionBot/Controller/MessagePushAPI/SNSAPI"
	"github.com/nyancatda/ShionBot/Controller/MessagePushAPI/SNSAPI/KaiHeiLaAPI"
	"github.com/nyancatda/ShionBot/Controller/MessagePushAPI/SNSAPI/LineAPI"
	"github.com/nyancatda/ShionBot/Controller/MessagePushAPI/SNSAPI/QQAPI"
	"github.com/nyancatda/ShionBot/Controller/MessagePushAPI/SNSAPI/TelegramAPI"
	"github.com/nyancatda/ShionBot/Utils/Language"
	"github.com/nyancatda/ShionBot/Utils/ViolationWordFilter"
)

/**
 * @description: 发送消息
 * @param {string} SNSName 聊天软件
 * @param {string} ChatType 聊天类型(Default,Friend,Group,GroupAt,Temp)
 * @param {string} UserID 接收消息的用户ID
 * @param {string} target 接受的聊天的ID
 * @param {string} text 消息文本
 * @param {bool} quote 是否需要回复
 * @param {string} quoteID 回复的消息ID(不需要时为空即可)
 * @param {string} AtID 需要@的人的ID(不需要时为空即可)
 * @param {int} group 临时会话群号(不需要时为0即可)
 * @return {[]byte}
 * @return {*http.Response}
 * @return {error}
 */
func SendMessage(SNSName string, ChatType string, UserID string, target string, text string, quote bool, quoteID string, AtID string, group int) ([]byte, *http.Response, error) {
	//判断违禁词
	if !ViolationWordFilter.ViolationWordFilter(text) {
		text = Language.Message(SNSName, UserID).ContainsIllegalWords
	}

	var Body []byte
	var HttpResponse *http.Response
	var err error

	switch SNSName {
	case "QQ":
		targets, _ := strconv.Atoi(target)
		switch ChatType {
		case "Friend":
			quoteIDs, _ := strconv.Atoi(quoteID)
			Body, HttpResponse, err = QQAPI.SendFriendMessage(targets, text, quote, quoteIDs)
		case "Group":
			quoteIDs, _ := strconv.Atoi(quoteID)
			Body, HttpResponse, err = QQAPI.SendGroupMessage(targets, text, quote, quoteIDs)
		case "GroupAt":
			AtID, _ := strconv.Atoi(AtID)
			quoteIDs, _ := strconv.Atoi(quoteID)
			Body, HttpResponse, err = QQAPI.SendGroupAtMessage(targets, text, AtID, quote, quoteIDs)
		case "Temp":
			quoteIDs, _ := strconv.Atoi(quoteID)
			Body, HttpResponse, err = QQAPI.SendTempMessage(targets, group, text, quote, quoteIDs)
		}
	case "Telegram":
		targets, _ := strconv.Atoi(target)
		switch ChatType {
		case "GroupAt":
			text = "@" + AtID + " " + text
			Body, HttpResponse, err = TelegramAPI.SendMessage(targets, text, true, false, 0, false)
		case "Group":
			quoteIDs, _ := strconv.Atoi(quoteID)
			Body, HttpResponse, err = TelegramAPI.SendMessage(targets, text, true, false, quoteIDs, quote)
		default:
			quoteIDs, _ := strconv.Atoi(quoteID)
			Body, HttpResponse, err = TelegramAPI.SendMessage(targets, text, true, false, quoteIDs, quote)
		}
	case "Line":
		switch ChatType {
		case "GroupAt":
			text = "@" + AtID + " " + text
			Body, HttpResponse, err = LineAPI.SendPushMessage(target, text, false)
		case "Group":
			if quote {
				Body, HttpResponse, err = LineAPI.SendReplyMessage(quoteID, text, false)
			} else {
				Body, HttpResponse, err = LineAPI.SendPushMessage(target, text, false)
			}
		default:
			Body, HttpResponse, err = LineAPI.SendPushMessage(target, text, false)
		}
	case "KaiHeiLa":
		switch ChatType {
		case "Group":
			Body, HttpResponse, err = KaiHeiLaAPI.SendChannelMessage(1, target, text, quote, quoteID)
		case "Friend":
			Body, HttpResponse, err = KaiHeiLaAPI.SendDirectMessage(1, target, text, quote, quoteID)
		}
	}

	//日志记录
	if len(Body) != 0 && err == nil {
		SNSAPI.Log(SNSName, ChatType, target, text)
	}

	return Body, HttpResponse, err
}

/**
 * @description: (QQ)发送头像戳一戳
 * @param {int} target 目标QQ号
 * @param {int} subject 消息接受主体，为群号或QQ号
 * @param {string} kind 上下文类型,可选值 Friend,Group,Stranger
 * @return {[]byte}
 * @return {*http.Response}
 * @return {error}
 */
func SendNudge(target int, subject int, kind string) ([]byte, *http.Response, error) {
	Body, HttpResponse, err := QQAPI.SendNudge(target, subject, kind)

	//日志记录
	if len(Body) != 0 && err == nil {
		SNSAPI.Log("QQ", kind, strconv.Itoa(subject), Language.DefaultLanguageMessage().Nudge+strconv.Itoa(target))
	}

	return Body, HttpResponse, err
}
