/*
 * @Author: NyanCatda
 * @Date: 2021-10-03 02:14:31
 * @LastEditTime: 2022-01-24 19:22:51
 * @LastEditors: NyanCatda
 * @Description: 获取Wiki页面信息
 * @FilePath: \ShionBot\src\Modular\GetWikiInfo\GetWikiInfo.go
 */
package GetWikiInfo

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/nyancatda/ShionBot/src/MediaWikiAPI"
	"github.com/nyancatda/ShionBot/src/Modular"
	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/Utils"
	"github.com/nyancatda/ShionBot/src/Utils/Language"
)

/**
 * @description: 错误处理
 * @param {string} SNSName
 * @param {string} UserID
 * @param {string} WikiLink
 * @param {string} title
 * @param {*Language.LanguageInfo} LanguageMessage
 * @return {*}
 */
func Error(SNSName string, UserID string, WikiLink string, title string, LanguageMessage *Language.LanguageInfo) string {
	text := Utils.StringVariable(LanguageMessage.GetWikiInfoError, []string{WikiLink, title})
	return text
}

/**
 * @description: 判断Wiki名字是否存在
 * @param {string} WikiName
 * @param {string} SNSName
 * @param {Struct.WebHookJson} Messagejson
 * @return {*}
 */
func WikiNameExist(WikiName string, SNSName string, Messagejson Struct.WebHookJson) bool {
	//判断用户设置
	db := Utils.SQLLiteLink()
	var user Struct.UserInfo
	UserID := Modular.GetSNSUserID(SNSName, Messagejson)
	db.Where("account = ? and sns_name = ?", UserID, SNSName).Find(&user)
	if user.Account == UserID {
		WikiInfo := user.WikiInfo
		WikiInfoData := []interface{}{}
		json.Unmarshal([]byte(WikiInfo), &WikiInfoData)
		for _, value := range WikiInfoData {
			WikiInfoName := value.(map[string]interface{})["WikiName"].(string)
			if find := strings.Contains(WikiName, WikiInfoName); find {
				return true
			}
		}
	}

	Config := Utils.GetConfig
	var ConfigWikiName string
	for one := range Config.Wiki.([]interface{}) {
		ConfigWikiName = Config.Wiki.([]interface{})[one].(map[interface{}]interface{})["WikiName"].(string)
		if find := strings.Contains(WikiName, ConfigWikiName); find {
			return true
		}
	}
	return false
}

/**
 * @description: 获取主Wiki名字
 * @param {string} SNSName
 * @param {Struct.WebHookJson} Messagejson
 * @return {*}
 */
func GeiMainWikiName(SNSName string, Messagejson Struct.WebHookJson) string {
	//获取用户设置
	db := Utils.SQLLiteLink()
	var user Struct.UserInfo
	UserID := Modular.GetSNSUserID(SNSName, Messagejson)
	db.Where("account = ? and sns_name = ?", UserID, SNSName).Find(&user)
	if user.Account == UserID {
		WikiInfo := user.WikiInfo
		WikiInfoData := []interface{}{}
		json.Unmarshal([]byte(WikiInfo), &WikiInfoData)
		for _, value := range WikiInfoData {
			WikiInfoName := value.(map[string]interface{})["WikiName"].(string)
			return WikiInfoName
		}
	}

	Config := Utils.GetConfig
	MainWikiName := Config.Wiki.([]interface{})[0].(map[interface{}]interface{})["WikiName"].(string)
	return MainWikiName
}

/**
 * @description: 搜索wiki
 * @param {string} SNSName
 * @param {Struct.WebHookJson} Messagejson
 * @param {string} WikiName
 * @param {string} title
 * @return {*}
 */
func SearchWiki(SNSName string, Messagejson Struct.WebHookJson, WikiName string, title string) string {
	WikiLink := Utils.GetWikiLink(SNSName, Messagejson, WikiName)
	SearchInfo, _ := MediaWikiAPI.Opensearch(WikiLink, 10, title)
	if len(SearchInfo) != 0 {
		SearchList := SearchInfo[1].([]interface{})
		if len(SearchList) != 0 {
			var SearchPages strings.Builder
			for _, value := range SearchList {
				PagseName := "[" + value.(string) + "]"
				SearchPages.WriteString(PagseName)
				SearchPages.WriteString("\n")
			}
			return SearchPages.String()
		}
		return ""
	} else {
		return ""
	}
}

/**
 * @description: 为空处理
 * @param {string} SNSName
 * @param {Struct.WebHookJson} Messagejson
 * @param {string} UserID
 * @param {string} WikiName
 * @param {string} title
 * @param {*Language.LanguageInfo} LanguageMessage
 * @return {*}
 */
func NilProcessing(SNSName string, Messagejson Struct.WebHookJson, UserID string, WikiName string, title string, LanguageMessage *Language.LanguageInfo) string {
	SearchInfo := SearchWiki(SNSName, Messagejson, WikiName, title)
	if SearchInfo != "" {
		Info := Utils.StringVariable(LanguageMessage.WikiInfoSearch, []string{SearchInfo, WikiName})
		return Info
	} else {
		WikiLink := Utils.GetWikiLink(SNSName, Messagejson, WikiName)
		return Error(SNSName, UserID, WikiLink, title, LanguageMessage)
	}
}

/**
 * @description: 获取Wiki页面标题，过滤后缀
 * @param {string} SNSName
 * @param {Struct.WebHookJson} Messagejson
 * @param {string} WikiName
 * @param {string} PageName
 * @return {*}
 */
func GetUrlTitle(SNSName string, Messagejson Struct.WebHookJson, WikiName string, PageName string) string {
	WikiLink := Utils.GetWikiLink(SNSName, Messagejson, WikiName)
	doc, err := htmlquery.LoadURL(WikiLink + "/" + PageName)
	if err != nil {
		fmt.Println(err)
	}
	for _, n := range htmlquery.Find(doc, "/html/head/title") {
		PageTitle := htmlquery.OutputHTML(n, false)
		countSplit := strings.SplitN(PageTitle, " - ", 2)
		Title := countSplit[0]
		return Title
	}
	return ""
}

/**
 * @description: 查询页面是否存在重定向
 * @param {string} SNSName
 * @param {Struct.WebHookJson} Messagejson
 * @param {string} WikiName
 * @param {string} title
 * @return {*}
 */
func QueryRedirects(SNSName string, Messagejson Struct.WebHookJson, WikiName string, title string) (whether bool, to string, from string, err error) {
	WikiLink := Utils.GetWikiLink(SNSName, Messagejson, WikiName)
	info, err := MediaWikiAPI.QueryRedirects(WikiLink, title)

	for _, value := range info.Query.Pages {
		if value.Title != "" {
			if len(info.Query.Normalized) != 0 {
				return true, info.Query.Normalized[0].To, info.Query.Normalized[0].From, err
			} else {
				PageTitleInfo := GetUrlTitle(SNSName, Messagejson, WikiName, title)
				if PageTitleInfo != title {
					ToTitle := PageTitleInfo
					return true, ToTitle, title, err
				}
			}
			return false, "", "", err
		}
	}

	return false, "", "", err
}

/**
 * @description: 获取Wiki页面信息
 * @param {string} SNSName
 * @param {Struct.WebHookJson} Messagejson
 * @param {string} UserID
 * @param {string} WikiName
 * @param {string} title
 * @param {string} language
 * @return {*}
 */
func GetWikiInfo(SNSName string, Messagejson Struct.WebHookJson, UserID string, WikiName string, title string, language string) (string, error) {
	var LanguageMessage *Language.LanguageInfo
	if language != "" {
		LanguageMessage = Language.DesignateLanguageMessage(language)
	} else {
		LanguageMessage = Language.Message(SNSName, UserID)
	}
	var err error
	RedirectsState, ToTitle, FromTitle, _ := QueryRedirects(SNSName, Messagejson, WikiName, title)
	var info MediaWikiAPI.QueryExtractsJson
	WikiLink := Utils.GetWikiLink(SNSName, Messagejson, WikiName)
	if RedirectsState {
		info, err = MediaWikiAPI.QueryExtracts(WikiLink, 100, ToTitle)
	} else {
		info, err = MediaWikiAPI.QueryExtracts(WikiLink, 100, title)
	}

	if len(info.Query.Pages) == 0 {
		return NilProcessing(SNSName, Messagejson, UserID, WikiName, title, LanguageMessage), err
	}

	var PageId string
	for one := range info.Query.Pages {
		PageId = one
	}

	if PageId != "-1" {
		PagesExtract := info.Query.Pages[PageId].Extract

		WikiLink := Utils.GetWikiLink(SNSName, Messagejson, WikiName)
		WikiPageInfo, err := MediaWikiAPI.QueryInfoUrl(WikiLink, title)
		var WikiPageLink string
		for _, Value := range WikiPageInfo.Query.Pages {
			WikiPageLink = Value.FullURL
		}

		var returnText string
		if RedirectsState {
			info := Utils.StringVariable(LanguageMessage.WikiInfoRedirect, []string{FromTitle, ToTitle})
			returnText = WikiPageLink + info + PagesExtract
		} else {
			returnText = WikiPageLink + "\n[" + title + "]\n" + PagesExtract
		}
		return returnText, err
	} else {
		return NilProcessing(SNSName, Messagejson, UserID, WikiName, title, LanguageMessage), err
	}
}
