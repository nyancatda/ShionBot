package QQAPI

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nyancatda/HttpRequest"
	"github.com/nyancatda/ShionBot/Utils/Language"
	"github.com/nyancatda/ShionBot/Utils/ReadConfig"
)

type returnJson struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	MessageId string `json:"messageId"`
}

/**
 * @description: 发送错误处理
 * @param {string} url 请求的URL
 * @param {map[string]interface{}} MessageBody 请求的信息
 * @return {[]byte}
 * @return {*http.Response}
 * @return {error}
 */
func sendError(url string, MessageBody map[string]interface{}) ([]byte, *http.Response, error) {
	//尝试通过创建一个新的SessionKey来从错误中恢复
	SessionKey, _, resp, err := CreateSessionKey()
	if err != nil {
		//无法获取SessionKey
		fmt.Println(Language.DefaultLanguageMessage().CannotConnectMirai)
		fmt.Println(err)
	} else if resp.Status != "200 OK" {
		//无法获取SessionKey
		fmt.Println(Language.DefaultLanguageMessage().CannotConnectMirai)
	} else {
		//将新的SessionKey写入消息链并重新发送
		MessageBody["sessionKey"] = SessionKey
		Body, _ := json.Marshal(MessageBody)
		Body, HttpResponse, err := HttpRequest.PostRequestJson(url, []string{}, string(Body))
		return Body, HttpResponse, err
	}
	var HttpResponse *http.Response
	return []byte{}, HttpResponse, err
}

/**
 * @description: 发送群消息
 * @param {int} target 群号
 * @param {string} text 消息文本
 * @param {bool} quote 是否需要回复
 * @param {int} quoteID 回复的消息ID(不需要时为0即可)
 * @return {[]byte}
 * @return {*http.Response}
 * @return {error}
 */
func SendGroupMessage(target int, text string, quote bool, quoteID int) ([]byte, *http.Response, error) {
	Config := ReadConfig.GetConfig
	sessionKey := GetSessionKey()

	//组成消息链
	MessageChain := make([]map[string]string, 1)
	MessageChain[0] = map[string]string{
		"type": "Plain",
		"text": text,
	}
	MessageBody := map[string]interface{}{
		"sessionKey":   sessionKey,
		"target":       target,
		"messageChain": MessageChain,
	}

	//判断是否需要引用回复
	if quote {
		MessageBody["quote"] = quoteID
	}

	MessageBodyJson, _ := json.Marshal(MessageBody)

	url := Config.SNS.QQ.APILink + "/sendGroupMessage"
	Body, HttpResponse, err := HttpRequest.PostRequestJson(url, []string{}, string(MessageBodyJson))

	//如果发送失败，则调用错误处理函数
	var config returnJson
	json.Unmarshal([]byte(Body), &config)
	if config.Code != 0 {
		return sendError(url, MessageBody)
	}

	return Body, HttpResponse, err
}

/**
 * @description: 发送带@的群消息
 * @param {int} target 群号
 * @param {string} text 消息文本
 * @param {int} AtID 需要@的人的QQ号
 * @param {bool} quote 是否需要回复
 * @param {int} quoteID 回复的消息ID(不需要时为0即可)
 * @return {[]byte}
 * @return {*http.Response}
 * @return {error}
 */
func SendGroupAtMessage(target int, text string, AtID int, quote bool, quoteID int) ([]byte, *http.Response, error) {
	Config := ReadConfig.GetConfig
	sessionKey := GetSessionKey()

	//组成消息链
	MessageChain := make([]map[string]interface{}, 2)
	MessageChain[0] = map[string]interface{}{
		"type":   "At",
		"target": AtID,
	}
	MessageChain[1] = map[string]interface{}{
		"type": "Plain",
		"text": text,
	}
	MessageBody := map[string]interface{}{
		"sessionKey":   sessionKey,
		"target":       target,
		"messageChain": MessageChain,
	}

	//判断是否需要引用回复
	if quote {
		MessageBody["quote"] = quoteID
	}

	MessageBodyJson, _ := json.Marshal(MessageBody)

	url := Config.SNS.QQ.APILink + "/sendGroupMessage"
	Body, HttpResponse, err := HttpRequest.PostRequestJson(url, []string{}, string(MessageBodyJson))

	//如果发送失败，则调用错误处理函数
	var config returnJson
	json.Unmarshal([]byte(Body), &config)
	if config.Code != 0 {
		return sendError(url, MessageBody)
	}

	return Body, HttpResponse, err
}

/**
 * @description: 发送头像戳一戳
 * @param {int} target 目标QQ号
 * @param {int} subject 消息接受主体，为群号或QQ号
 * @param {string} kind 上下文类型,可选值 Friend,Group,Stranger
 * @return {[]byte}
 * @return {*http.Response}
 * @return {error}
 */
func SendNudge(target int, subject int, kind string) ([]byte, *http.Response, error) {
	Config := ReadConfig.GetConfig
	sessionKey := GetSessionKey()

	//组成消息链
	MessageBody := map[string]interface{}{
		"sessionKey": sessionKey,
		"target":     target,
		"subject":    subject,
		"kind":       kind,
	}

	MessageBodyJson, _ := json.Marshal(MessageBody)

	url := Config.SNS.QQ.APILink + "/sendNudge"
	Body, HttpResponse, err := HttpRequest.PostRequestJson(url, []string{}, string(MessageBodyJson))

	//如果发送失败，则调用错误处理函数
	var config returnJson
	json.Unmarshal([]byte(Body), &config)
	if config.Code != 0 {
		return sendError(url, MessageBody)
	}

	return Body, HttpResponse, err
}

/**
 * @description: 发送好友消息
 * @param {int} target 好友QQ号
 * @param {string} text 消息文本
 * @param {bool} quote 是否需要回复
 * @param {int} quoteID 回复的消息ID(不需要时为0即可)
 * @return {[]byte}
 * @return {*http.Response}
 * @return {error}
 */
func SendFriendMessage(target int, text string, quote bool, quoteID int) ([]byte, *http.Response, error) {
	Config := ReadConfig.GetConfig
	sessionKey := GetSessionKey()

	//组成消息链
	MessageChain := make([]map[string]string, 1)
	MessageChain[0] = map[string]string{
		"type": "Plain",
		"text": text,
	}
	MessageBody := map[string]interface{}{
		"sessionKey":   sessionKey,
		"target":       target,
		"messageChain": MessageChain,
	}

	//判断是否需要引用回复
	if quote {
		MessageBody["quote"] = quoteID
	}

	MessageBodyJson, _ := json.Marshal(MessageBody)

	url := Config.SNS.QQ.APILink + "/sendFriendMessage"
	Body, HttpResponse, err := HttpRequest.PostRequestJson(url, []string{}, string(MessageBodyJson))

	//如果发送失败，则调用错误处理函数
	var config returnJson
	json.Unmarshal([]byte(Body), &config)
	if config.Code != 0 {
		return sendError(url, MessageBody)
	}

	return Body, HttpResponse, err
}

/**
 * @description: 发送临时会话
 * @param {int} target 临时会话对象QQ号
 * @param {int} group 临时会话群号
 * @param {string} text 消息文本
 * @param {bool} quote 是否需要回复
 * @param {int} quoteID 回复的消息ID(不需要时为0即可)
 * @return {[]byte}
 * @return {*http.Response}
 * @return {error}
 */
func SendTempMessage(target int, group int, text string, quote bool, quoteID int) ([]byte, *http.Response, error) {
	Config := ReadConfig.GetConfig
	sessionKey := GetSessionKey()

	//组成消息链
	MessageChain := make([]map[string]string, 1)
	MessageChain[0] = map[string]string{
		"type": "Plain",
		"text": text,
	}
	MessageBody := map[string]interface{}{
		"sessionKey":   sessionKey,
		"qq":           target,
		"group":        group,
		"messageChain": MessageChain,
	}

	//判断是否需要引用回复
	if quote {
		MessageBody["quote"] = quoteID
	}

	MessageBodyJson, _ := json.Marshal(MessageBody)

	url := Config.SNS.QQ.APILink + "/sendTempMessage"
	Body, HttpResponse, err := HttpRequest.PostRequestJson(url, []string{}, string(MessageBodyJson))

	//如果发送失败，则调用错误处理函数
	var config returnJson
	json.Unmarshal([]byte(Body), &config)
	if config.Code != 0 {
		return sendError(url, MessageBody)
	}

	return Body, HttpResponse, err
}
