package Plugin

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"strings"
	"xyz.nyan/MediaWiki-Bot/MediaWikiAPI"
)

func Error(WikiLink string, title string) string {
	return "在[" + WikiLink + "]中找不到[" + title + "]哦，请检查输入是否正确"
}

//获取Wiki页面标题，过滤后缀
func GetUrlTitle(WikiName string, PageName string) string {
	WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
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

	_, ok := info["query"]
	if ok {
		normalized, ok := info["query"].(map[string]interface{})["normalized"]
		if ok {
			return true, normalized.([]interface{})[0].(map[string]interface{})["to"].(string), normalized.([]interface{})[0].(map[string]interface{})["from"].(string)
		} else {
			PageTitleInfo := GetUrlTitle(WikiName, title)
			if PageTitleInfo != title {
				ToTitle := PageTitleInfo
				return true, ToTitle, title
			}
		}
		return false, "", ""
	}
	return false, "", ""
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

	_, ok := info["query"]
	if !ok {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		return Error(WikiLink,title)
	}

	pagesIdInfo, ok := info["query"].(map[string]interface{})["pages"]
	if !ok {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		return Error(WikiLink,title)
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
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		return Error(WikiLink,title)
	}
}
