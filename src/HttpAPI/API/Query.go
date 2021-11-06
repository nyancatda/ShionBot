package API

import (
	"github.com/gin-gonic/gin"
	"xyz.nyan/MediaWiki-Bot/src/HttpAPI/JsonProcessing"
	"xyz.nyan/MediaWiki-Bot/src/MediaWikiAPI"
	"xyz.nyan/MediaWiki-Bot/src/Plugin"
	"xyz.nyan/MediaWiki-Bot/src/utils"
	"xyz.nyan/MediaWiki-Bot/src/utils/Language"
)

func QueryInfo(c *gin.Context) map[string]interface{} {
	title := c.DefaultQuery("title", "")
	Config := utils.ReadConfig()
	MainWikiName := Config.Wiki.([]interface{})[0].(map[interface{}]interface{})["WikiName"].(string)
	WikiName := c.DefaultQuery("wiki_name", MainWikiName)
	WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
	language := c.DefaultQuery("language", "zh-CN")
	if !Language.LanguageExist(language) {
		JsonInfo := JsonProcessing.JsonRoot(5000, Language.StringVariable(1, Language.DefaultLanguageMessage().LanguageModificationFailed, language, ""))
		return JsonInfo
	}
	if title == "" {
		JsonInfo := JsonProcessing.JsonRoot(5000, Language.DefaultLanguageMessage().TitleEmpty)
		return JsonInfo
	}

	WikiInfo, err := Plugin.GetWikiInfo("HttpAPI", "HttpAPI", WikiName, title, language)
	if err != nil {
		JsonInfo := JsonProcessing.JsonRoot(5000, Language.StringVariable(1, Language.DefaultLanguageMessage().WikiLinkError, WikiLink, ""))
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
