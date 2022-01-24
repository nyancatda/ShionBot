/*
 * @Author: NyanCatda
 * @Date: 2021-11-06 21:49:39
 * @LastEditTime: 2022-01-24 18:07:11
 * @LastEditors: NyanCatda
 * @Description: 查询Wiki信息
 * @FilePath: \ShionBot\src\HttpAPI\API\Query.go
 */
package API

import (
	"github.com/gin-gonic/gin"
	"github.com/nyancatda/ShionBot/src/HttpAPI/JsonProcessing"
	"github.com/nyancatda/ShionBot/src/Modular/GetWikiInfo"
	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/Utils"
	"github.com/nyancatda/ShionBot/src/Utils/Language"
)

func QueryInfo(c *gin.Context) map[string]interface{} {
	title := c.DefaultQuery("title", "")
	Config := Utils.GetConfig
	MainWikiName := Config.Wiki.([]interface{})[0].(map[interface{}]interface{})["WikiName"].(string)
	WikiName := c.DefaultQuery("wiki_name", MainWikiName)
	var Messagejson Struct.WebHookJson
	WikiLink := Utils.GetWikiLink("HttpAPI", Messagejson, WikiName)
	language := c.DefaultQuery("language", "zh-CN")
	if !Language.LanguageExist(language) {
		JsonInfo := JsonProcessing.JsonRoot(5000, Utils.StringVariable(Language.DefaultLanguageMessage().LanguageModificationFailed, []string{language}))
		return JsonInfo
	}
	if title == "" {
		JsonInfo := JsonProcessing.JsonRoot(5000, Language.DefaultLanguageMessage().TitleEmpty)
		return JsonInfo
	}

	WikiInfo, err := GetWikiInfo.GetWikiInfo("HttpAPI", Messagejson, "HttpAPI", WikiName, title, language)
	if err != nil {
		JsonInfo := JsonProcessing.JsonRoot(5000, Utils.StringVariable(Language.DefaultLanguageMessage().WikiLinkError, []string{WikiLink}))
		return JsonInfo
	}

	JsonInfo := JsonProcessing.JsonRoot(200, "")
	JsonInfo["data"] = map[string]interface{}{
		"wiki_name": WikiName,
		"wiki_link": WikiLink,
		"text":      WikiInfo,
	}

	return JsonInfo
}
