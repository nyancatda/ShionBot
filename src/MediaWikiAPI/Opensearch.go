/*
 * @Author: NyanCatda
 * @Date: 2021-10-12 16:10:33
 * @LastEditTime: 2022-01-24 17:49:47
 * @LastEditors: NyanCatda
 * @Description: MediaWiki OpensearchAPI封装
 * @FilePath: \ShionBot\src\MediaWikiAPI\Opensearch.go
 */
package MediaWikiAPI

import (
	"encoding/json"

	"github.com/nyancatda/ShionBot/src/utils"
)

/**
 * @description: 使用开放搜索协议搜索wiki(最多返回10条)，通常来说[0]为搜索标题，[1]为条目标题列表，[2]不知道，[3]为条目链接列表，对应[1]
 * @param {string} WikiLink Wiki链接
 * @param {string} title 需要搜索的页面标题
 * @return {*}
 */
func Opensearch(WikiLink string, title string) ([]string, error) {
	url := WikiLink + "/api.php?action=opensearch&limit=10&redirects=resolve&search=" + title
	body, err := utils.HttpRequest(url)

	var info []string
	json.Unmarshal([]byte(body), &info)
	return info, err
}
