//MediaWikiAPI封装
package MediaWikiAPI

import (
	"encoding/json"

	"xyz.nyan/MediaWiki-Bot/utils"
	"strconv"
)

//从配置文件读取Wiki名字对应的Wiki链接
func GetWikiLink(WikiName string) (string) {
	Config := utils.ReadConfig()
	Map := Config.Wiki
	var ConfigWikiName string
	for one := range Map {
		ConfigWikiName = one
		if ConfigWikiName == WikiName {
			return Config.Wiki[ConfigWikiName].(map[interface{}]interface{})["WikiLink"].(string)
		}
	}
	return ""
}

//查询页面信息，返回带URL
//WikiName 需要查询的Wiki名字
//title 需要查询的页面标题
func QueryInfoUrl(WikiName string,title string) (map[string]interface{}) {
	WikiLink := GetWikiLink(WikiName)
	url := WikiLink + "/api.php?action=query&prop=info&inprop=url&format=json&titles=" + title
	body := utils.HttpRequest(url)
	
    info := make(map[string]interface{})
	json.Unmarshal([]byte(body), &info)
	return info
}

//查询页面重定向信息
//WikiName 需要查询的Wiki名字
//title 需要查询的页面标题
func QueryRedirects(WikiName string,title string) (map[string]interface{}) {
	WikiLink := GetWikiLink(WikiName)
    url := WikiLink+"/api.php?action=query&prop=redirects&format=json&titles="+title
    body := utils.HttpRequest(url)

	info := make(map[string]interface{})
	json.Unmarshal([]byte(body), &info)
	return info
}

//查询页面内容提取物
//WikiName 需要查询的Wiki名字
//exchars 返回的字数
//title 需要查询的页面标题
func QueryExtracts(WikiName string,exchars int,title string) (map[string]interface{}) {
	WikiLink := GetWikiLink(WikiName)
	url := WikiLink+"/api.php?action=query&prop=extracts&exchars="+strconv.Itoa(exchars)+"&explaintext=true&format=json&titles="+title
	body := utils.HttpRequest(url)

	info := make(map[string]interface{})
	json.Unmarshal([]byte(body), &info)
	return info
}

//查询页面修订信息
//WikiName 需要查询的Wiki名字
//title 需要查询的页面标题
func QueryRevisions(WikiName string,title string) (map[string]interface{}) {
	WikiLink := GetWikiLink(WikiName)
    url := WikiLink+"/api.php?action=query&prop=revisions&format=json&titles="+title
    body := utils.HttpRequest(url)

	info := make(map[string]interface{})
	json.Unmarshal([]byte(body), &info)
	return info
}