//MediaWikiAPI封装
package MediaWikiAPI

import (
	"encoding/json"
	"strconv"

	"github.com/nyancatda/ShionBot/src/Modular"
	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/utils"
)

//读取Wiki名字对应的Wiki链接
func GetWikiLink(SNSName string, Messagejson Struct.WebHookJson, WikiName string) string {
	//获取用户配置
	db := utils.SQLLiteLink()
	var user Struct.UserInfo
	UserID := Modular.GetSNSUserID(SNSName, Messagejson)
	db.Where("account = ? and sns_name = ?", UserID, SNSName).Find(&user)
	if user.Account == UserID {
		WikiInfo := user.WikiInfo
		WikiInfoData := []interface{}{}
		json.Unmarshal([]byte(WikiInfo), &WikiInfoData)
		for _, value := range WikiInfoData {
			WikiInfoName := value.(map[string]interface{})["WikiName"].(string)
			if WikiName == WikiInfoName {
				return "https://" + value.(map[string]interface{})["WikiLink"].(string)
			}
		}
	}

	Config := utils.GetConfig
	var ConfigWikiName string
	for one := range Config.Wiki.([]interface{}) {
		ConfigWikiName = Config.Wiki.([]interface{})[one].(map[interface{}]interface{})["WikiName"].(string)
		if ConfigWikiName == WikiName {
			return Config.Wiki.([]interface{})[one].(map[interface{}]interface{})["WikiLink"].(string)
		}
	}
	return ""
}

//查询页面信息，返回带URL
//SNSName 聊天软件名字
//Messagejson 消息json
//WikiName 需要查询的Wiki名字
//title 需要查询的页面标题
func QueryInfoUrl(SNSName string, Messagejson Struct.WebHookJson, WikiName string, title string) (map[string]interface{}, error) {
	WikiLink := GetWikiLink(SNSName, Messagejson, WikiName)
	url := WikiLink + "/api.php?action=query&prop=info&inprop=url&format=json&titles=" + title
	body, err := utils.HttpRequest(url)

	info := make(map[string]interface{})
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
