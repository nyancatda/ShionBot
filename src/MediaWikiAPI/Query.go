/*
 * @Author: NyanCatda
 * @Date: 2021-10-03 20:50:06
 * @LastEditTime: 2022-01-24 16:48:11
 * @LastEditors: NyanCatda
 * @Description: MediaWikiAPI封装
 * @FilePath: \ShionBot\src\MediaWikiAPI\Query.go
 */
package MediaWikiAPI

import (
	"encoding/json"
	"strconv"

	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/utils"
)

type QueryInfoUrlJson struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         struct {
		Pages map[string]struct {
			PageID               int    `json:"pageid"` //页面ID
			Ns                   int    `json:"ns"`
			Title                string `json:"title"`                //页面标题
			Contentmodel         string `json:"contentmodel"`         //页面类型
			Pagelanguage         string `json:"pagelanguage"`         //页面语言
			Pagelanguagehtmlcode string `json:"pagelanguagehtmlcode"` //页面HTML语言
			Pagelanguagedir      string `json:"pagelanguagedir"`
			Touched              string `json:"touched"` //创建时间
			Lastrevid            int    `json:"lastrevid"`
			Length               int    `json:"length"`
			FullURL              string `json:"fullurl"`      //页面完整URL
			EditURL              string `json:"editurl"`      //页面编辑URL
			CanonicalURL         string `json:"canonicalurl"` //页面规范的URL
		} `json:"pages"`
	} `json:"query"`
}

/**
 * @description: 查询页面信息，返回带URL
 * @param {string} WikiLink Wiki连接
 * @param {string} title 页面标题
 * @return {*}
 */
func QueryInfoUrl(WikiLink string, title string) (QueryInfoUrlJson, error) {
	url := WikiLink + "/api.php?action=query&prop=info&inprop=url&format=json&titles=" + title
	body, err := utils.HttpRequest(url)

	var info QueryInfoUrlJson
	json.Unmarshal([]byte(body), &info)
	return info, err
}

//查询页面重定向信息
//SNSName 聊天软件名字
//Messagejson 消息json
//WikiName 需要查询的Wiki名字
//title 需要查询的页面标题
func QueryRedirects(SNSName string, Messagejson Struct.WebHookJson, WikiName string, title string) (map[string]interface{}, error) {
	WikiLink := GetWikiLink(SNSName, Messagejson, WikiName)
	url := WikiLink + "/api.php?action=query&prop=redirects&format=json&titles=" + title
	body, err := utils.HttpRequest(url)

	info := make(map[string]interface{})
	json.Unmarshal([]byte(body), &info)
	return info, err
}

//查询页面内容提取物
//SNSName 聊天软件名字
//Messagejson 消息json
//WikiName 需要查询的Wiki名字
//exchars 返回的字数
//title 需要查询的页面标题
func QueryExtracts(SNSName string, Messagejson Struct.WebHookJson, WikiName string, exchars int, title string) (map[string]interface{}, error) {
	WikiLink := GetWikiLink(SNSName, Messagejson, WikiName)
	url := WikiLink + "/api.php?action=query&prop=extracts&exchars=" + strconv.Itoa(exchars) + "&explaintext=true&format=json&titles=" + title
	body, err := utils.HttpRequest(url)

	info := make(map[string]interface{})
	json.Unmarshal([]byte(body), &info)
	return info, err
}

//查询页面修订信息
//SNSName 聊天软件名字
//Messagejson 消息json
//WikiName 需要查询的Wiki名字
//title 需要查询的页面标题
func QueryRevisions(SNSName string, Messagejson Struct.WebHookJson, WikiName string, title string) (map[string]interface{}, error) {
	WikiLink := GetWikiLink(SNSName, Messagejson, WikiName)
	url := WikiLink + "/api.php?action=query&prop=revisions&format=json&titles=" + title
	body, err := utils.HttpRequest(url)

	info := make(map[string]interface{})
	json.Unmarshal([]byte(body), &info)
	return info, err
}

//查询网站的全部系统信息
//WikiLink Wiki链接
func QuerySiteinfoGeneral(WikiLink string) (map[string]interface{}, error) {
	url := WikiLink + "/api.php?action=query&meta=siteinfo&siprop=general&format=json"
	body, err := utils.HttpRequest(url)

	info := make(map[string]interface{})
	json.Unmarshal([]byte(body), &info)
	return info, err
}
