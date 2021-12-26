package API

import (
	"github.com/gin-gonic/gin"
	"xyz.nyan/ShionBot/src/HttpAPI/JsonProcessing"
	"xyz.nyan/ShionBot/src/MediaWikiAPI"
	"xyz.nyan/ShionBot/src/Plugin/GetWikiInfo"
	"xyz.nyan/ShionBot/src/Struct"
	"xyz.nyan/ShionBot/src/utils"
	"xyz.nyan/ShionBot/src/utils/Language"
)

func QueryInfo(c *gin.Context) map[string]interface{} {
	title := c.DefaultQuery("title", "")
	Config := utils.GetConfig
	MainWikiName := Config.Wiki.([]interface{})[0].(map[interface{}]interface{})["WikiName"].(string)
	WikiName := c.DefaultQuery("wiki_name", MainWikiName)
	var Messagejson Struct.WebHookJson
	WikiLink := MediaWikiAPI.GetWikiLink("HttpAPI", Messagejson, WikiName)
	language := c.DefaultQuery("language", "zh-CN")
	if !Language.LanguageExist(language) {
		JsonInfo := JsonProcessing.JsonRoot(5000, utils.StringVariable(Language.DefaultLanguageMessage().LanguageModificationFailed, []string{language}))
		return JsonInfo
	}
	if title == "" {
		JsonInfo := JsonProcessing.JsonRoot(5000, Language.DefaultLanguageMessage().TitleEmpty)
		return JsonInfo
	}

	WikiInfo, err := GetWikiInfo.GetWikiInfo("HttpAPI", Messagejson, "HttpAPI", WikiName, title, language)
	if err != nil {
		JsonInfo := JsonProcessing.JsonRoot(5000, utils.StringVariable(Language.DefaultLanguageMessage().WikiLinkError, []string{WikiLink}))
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
