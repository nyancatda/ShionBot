package utils

import (
	"encoding/json"
)

func Error(title string) (string) {
    return "找不到["+title+"]哦，请检查输入是否正确"
}

//获取Wiki页面信息
func QueryWikiInfo(title string) (interface{}) {
    Config := ReadConfig()
    url := Config.Wiki.WikiLink+"/api.php?action=query&prop=info&inprop=url&format=json&titles="+title
    body := HttpRequest(url)

    info := make(map[string]interface{})
	json.Unmarshal([]byte(body), &info)
    pagesIdInfo := info["query"].(map[string]interface{})["pages"]
    var PageId string
    for one := range pagesIdInfo.(map[string]interface{}) {
        PageId = one
    }

    return info["query"].(map[string]interface{})["pages"].(map[string]interface{})[PageId]
}

//查询页面是否存在重定向
func QueryRedirects(title string) (bool,interface{}) {
    Config := ReadConfig()
    url := Config.Wiki.WikiLink+"/api.php?action=query&prop=redirects&format=json&titles="+title
    body := HttpRequest(url)

	info := make(map[string]interface{})
	json.Unmarshal([]byte(body), &info)
    //pagesIdInfo := info["query"].(map[string]interface{})["pages"]

    if normalized, ok := info["query"].(map[string]interface{})["normalized"]; ok {
        return true,normalized
    } else {
        return false,info
    }
}

//获取Wiki页面信息
func GetWikiInfo(title string) (string) {
    Config := ReadConfig()
    RedirectsState,Redirectsinfo := QueryRedirects(title)
    var url,ToTitle,FromTitle string
    if RedirectsState {
        ToTitle = Redirectsinfo.([]interface{})[0].(map[string]interface{})["to"].(string)
        FromTitle = Redirectsinfo.([]interface{})[0].(map[string]interface{})["from"].(string)
        url = Config.Wiki.WikiLink+"/api.php?action=query&prop=extracts&exchars=100&explaintext=true&format=json&titles="+ToTitle
    } else {
        url = Config.Wiki.WikiLink+"/api.php?action=query&prop=extracts&exchars=100&explaintext=true&format=json&titles="+title
    }

    body := HttpRequest(url)

	info := make(map[string]interface{})
	json.Unmarshal([]byte(body), &info)

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
            WikiPageInfo := QueryWikiInfo(ToTitle)
            WikiPageLink := WikiPageInfo.(map[string]interface{})["fullurl"].(string)
            returnText = WikiPageLink+"\n(重定向["+FromTitle+"]->["+ToTitle+"])\n"+PagesExtract.(string)
        } else {
            WikiPageInfo := QueryWikiInfo(title)
            WikiPageLink := WikiPageInfo.(map[string]interface{})["fullurl"].(string)
            returnText = WikiPageLink+"\n["+title+"]\n"+PagesExtract.(string)
        }
        return returnText
    } else {
        return Error(title)
    }
}
