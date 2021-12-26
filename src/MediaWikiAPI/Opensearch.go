package MediaWikiAPI

import (
	"encoding/json"

	"github.com/nyancatda/ShionBot/src/Struct"
	"github.com/nyancatda/ShionBot/src/utils"
)

//使用开放搜索协议搜索wiki(最多返回10条)
//SNSName 聊天软件名字
//Messagejson 消息json
//WikiName 需要查询的Wiki名字
//title 需要搜索的页面标题
func Opensearch(SNSName string, Messagejson Struct.WebHookJson, WikiName string, title string) ([]interface{}, error) {
	WikiLink := GetWikiLink(SNSName, Messagejson, WikiName)
	url := WikiLink + "/api.php?action=opensearch&limit=10&redirects=resolve&search=" + title
	body, err := utils.HttpRequest(url)

	info := []interface{}{}
	json.Unmarshal([]byte(body), &info)
	return info, err
}
