//MediaWikiAPI封装
package MediaWikiAPI

import (
	"encoding/json"

	"xyz.nyan/MediaWiki-Bot/utils"
	"strconv"
)

//查询页面信息，返回带URL
//title 需要查询的页面标题
func QueryInfoUrl(title string) (map[string]interface{}) {
	Config := utils.ReadConfig()
	url := Config.Wiki.WikiLink + "/api.php?action=query&prop=info&inprop=url&format=json&titles=" + title
	body := utils.HttpRequest(url)
	
    info := make(map[string]interface{})
	json.Unmarshal([]byte(body), &info)
	return info
}

//查询页面重定向信息
//title 需要查询的页面标题
func QueryRedirects(title string) (map[string]interface{}) {
	Config := utils.ReadConfig()
    url := Config.Wiki.WikiLink+"/api.php?action=query&prop=redirects&format=json&titles="+title
    body := utils.HttpRequest(url)

	info := make(map[string]interface{})
	json.Unmarshal([]byte(body), &info)
	return info
}

//查询页面内容提取物
//exchars 返回的字数
//title 需要查询的页面标题
func QueryExtracts(exchars int,title string) (map[string]interface{}) {
	Config := utils.ReadConfig()
	url := Config.Wiki.WikiLink+"/api.php?action=query&prop=extracts&exchars="+strconv.Itoa(exchars)+"&explaintext=true&format=json&titles="+title
	body := utils.HttpRequest(url)

	info := make(map[string]interface{})
	json.Unmarshal([]byte(body), &info)
	return info
}
