package MediaWikiAPI

import (
	"encoding/json"

	"xyz.nyan/MediaWiki-Bot/utils"
)

//使用开放搜索协议搜索wiki(最多返回10条)
//WikiName 需要查询的Wiki名字
//title 需要搜索的页面标题
func Opensearch(WikiName string, title string) ([]interface{}, error) {
	WikiLink := GetWikiLink(WikiName)
	url := WikiLink + "/api.php?action=opensearch&limit=10&redirects=resolve&search=" + title
	body, err := utils.HttpRequest(url)

	info := []interface{}{}
	json.Unmarshal([]byte(body), &info)
	return info, err
}
