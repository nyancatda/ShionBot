package Plugin

import (
	"fmt"
	"strings"

	"xyz.nyan/MediaWiki-Bot/MediaWikiAPI"
)

func Error(title string) string {
	return "找不到[" + title + "]哦，请检查输入是否正确"
}

//获取Wiki页面信息
func QueryWikiInfo(WikiName string, title string) interface{} {
	info := MediaWikiAPI.QueryInfoUrl(WikiName, title)
	pagesIdInfo := info["query"].(map[string]interface{})["pages"]
	var PageId string
	for one := range pagesIdInfo.(map[string]interface{}) {
		PageId = one
	}

	return info["query"].(map[string]interface{})["pages"].(map[string]interface{})[PageId]
}

//查询页面是否存在重定向
func QueryRedirects(WikiName string, title string) (whether bool, to string, from string) {
	info := MediaWikiAPI.QueryRedirects(WikiName, title)

	if normalized, ok := info["query"].(map[string]interface{})["normalized"]; ok {
		return true, normalized.([]interface{})[0].(map[string]interface{})["to"].(string), normalized.([]interface{})[0].(map[string]interface{})["from"].(string)
	} else {
		RevisionsInfo := MediaWikiAPI.QueryRevisions(WikiName, title)
		pagesIdInfo, ok := RevisionsInfo["query"].(map[string]interface{})["pages"]
		if ok {
			var PageId string
			for one := range pagesIdInfo.(map[string]interface{}) {
				PageId = one
			}
			commentInfo, ok := RevisionsInfo["query"].(map[string]interface{})["pages"].(map[string]interface{})[PageId].(map[string]interface{})["revisions"]
			if ok {
				find := strings.Contains(commentInfo.([]interface{})[0].(map[string]interface{})["comment"].(string), "重定向页面至")
				if find {
					trimStr := strings.Trim(commentInfo.([]interface{})[0].(map[string]interface{})["comment"].(string), "重定向页面至")
					trimStr = strings.Trim(trimStr, "[")
					ToTitle := strings.Trim(trimStr, "]")
					fmt.Println(ToTitle)
					return true, ToTitle, title
				}
			}
		}
		return false, "", ""
	}
}

//获取Wiki页面信息
func GetWikiInfo(WikiName string, title string) string {
	RedirectsState, ToTitle, FromTitle := QueryRedirects(WikiName, title)
	var info map[string]interface{}
	if RedirectsState {
		info = MediaWikiAPI.QueryExtracts(WikiName, 100, ToTitle)
	} else {
		info = MediaWikiAPI.QueryExtracts(WikiName, 100, title)
	}

	pagesIdInfo, ok := info["query"].(map[string]interface{})["pages"]
	if !ok {
		return Error(title)
	}

	var PageId string
	for one := range pagesIdInfo.(map[string]interface{}) {
		PageId = one
	}

	if PageId != "-1" {
		PagesExtract := info["query"].(map[string]interface{})["pages"].(map[string]interface{})[PageId].(map[string]interface{})["extract"]
		var returnText string
		if RedirectsState {
			WikiPageInfo := QueryWikiInfo(WikiName, ToTitle)
			WikiPageLink := WikiPageInfo.(map[string]interface{})["fullurl"].(string)
			returnText = WikiPageLink + "\n(重定向[" + FromTitle + "]->[" + ToTitle + "])\n" + PagesExtract.(string)
		} else {
			WikiPageInfo := QueryWikiInfo(WikiName, title)
			WikiPageLink := WikiPageInfo.(map[string]interface{})["fullurl"].(string)
			returnText = WikiPageLink + "\n[" + title + "]\n" + PagesExtract.(string)
		}
		return returnText
	} else {
		return Error(title)
	}
}
