/*
 * @Author: NyanCatda
 * @Date: 2021-10-12 16:10:33
 * @LastEditTime: 2022-01-24 18:06:32
 * @LastEditors: NyanCatda
 * @Description: MediaWiki OpensearchAPI封装
 * @FilePath: \ShionBot\src\MediaWikiAPI\Opensearch.go
 */
package MediaWikiAPI

import (
	"encoding/json"
	"strconv"

	"github.com/nyancatda/ShionBot/src/utils"
)

/**
 * @description: 使用开放搜索协议搜索wiki，通常来说[0]为搜索标题，[1]为条目标题列表，[2]不知道，[3]为条目链接列表，对应[1]
 * @param {string} WikiLink Wiki链接
 * @param {int} Limit 返回条数限制
 * @param {string} title 需要搜索的页面标题
 * @return {*}
 */
func Opensearch(WikiLink string, Limit int, title string) ([]interface{}, error) {
	url := WikiLink + "/api.php?action=opensearch&limit=" + strconv.Itoa(Limit) + "&redirects=resolve&search=" + title
	body, err := utils.HttpRequest(url)

	var info []interface{}
	json.Unmarshal([]byte(body), &info)
	return info, err
}
