/*
 * @Author: NyanCatda
 * @Date: 2021-10-03 20:50:06
 * @LastEditTime: 2022-01-24 19:27:01
 * @LastEditors: NyanCatda
 * @Description: MediaWiki查询类API封装
 * @FilePath: \ShionBot\src\MediaWikiAPI\Query.go
 */
package MediaWikiAPI

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"github.com/nyancatda/ShionBot/src/Utils"
)

type QueryInfoUrlJson struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         struct {
		Pages map[string]struct {
			PageID               int       `json:"pageid"` //页面ID
			Ns                   int       `json:"ns"`
			Title                string    `json:"title"`                //页面标题
			Contentmodel         string    `json:"contentmodel"`         //页面类型
			Pagelanguage         string    `json:"pagelanguage"`         //页面语言
			Pagelanguagehtmlcode string    `json:"pagelanguagehtmlcode"` //页面HTML语言
			Pagelanguagedir      string    `json:"pagelanguagedir"`
			Touched              time.Time `json:"touched"` //创建时间
			Lastrevid            int       `json:"lastrevid"`
			Length               int       `json:"length"`
			FullURL              string    `json:"fullurl"`      //页面完整URL
			EditURL              string    `json:"editurl"`      //页面编辑URL
			CanonicalURL         string    `json:"canonicalurl"` //页面规范的URL
		} `json:"pages"`
	} `json:"query"`
}

/**
 * @description: 查询页面信息，返回带URL
 * @param {string} WikiLink Wiki连接
 * @param {string} title 页面标题
 * @return {QueryInfoUrlJson}
 * @return {error}
 */
func QueryInfoUrl(WikiLink string, title string) (QueryInfoUrlJson, error) {
	title = url.QueryEscape(title)
	url := WikiLink + "/api.php?action=query&prop=info&inprop=url&format=json&titles=" + title
	body, err := Utils.HttpRequest(url)

	var info QueryInfoUrlJson
	json.Unmarshal([]byte(body), &info)
	return info, err
}

type QueryRedirectsJson struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         struct {
		Normalized []struct {
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"normalized"`
		Pages map[string]struct {
			Ns      int    `json:"ns"`
			Title   string `json:"title"` //页面标题
			Missing string `json:"missing"`
		} `json:"pages"`
	} `json:"query"`
}

/**
 * @description: 查询页面重定向信息
 * @param {string} WikiLink Wiki链接
 * @param {string} title 需要查询的页面标题
 * @return {QueryRedirectsJson}
 * @return {error}
 */
func QueryRedirects(WikiLink string, title string) (QueryRedirectsJson, error) {
	title = url.QueryEscape(title)
	url := WikiLink + "/api.php?action=query&prop=redirects&format=json&titles=" + title
	body, err := Utils.HttpRequest(url)

	var info QueryRedirectsJson
	json.Unmarshal([]byte(body), &info)
	return info, err
}

type QueryExtractsJson struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         struct {
		Pages map[string]struct {
			PageID  int    `json:"pageid"` //页面ID
			Ns      int    `json:"ns"`
			Title   string `json:"title"`   //页面标题
			Extract string `json:"extract"` //页面提取物
		} `json:"pages"`
	} `json:"query"`
}

/**
 * @description: 查询页面内容提取物
 * @param {string} WikiLink Wiki链接
 * @param {int} exchars 返回的字数
 * @param {string} title 需要查询的页面标题
 * @return {QueryExtractsJson}
 * @return {error}
 */
func QueryExtracts(WikiLink string, exchars int, title string) (QueryExtractsJson, error) {
	title = url.QueryEscape(title)
	url := WikiLink + "/api.php?action=query&prop=extracts&exchars=" + strconv.Itoa(exchars) + "&explaintext=true&format=json&titles=" + title
	body, err := Utils.HttpRequest(url)

	var info QueryExtractsJson
	json.Unmarshal([]byte(body), &info)
	return info, err
}

type QueryRevisionsJson struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         struct {
		Pages map[string]struct {
			PageID    int        `json:"pageid"` //页面ID
			Ns        int        `json:"ns"`
			Title     string     `json:"title"` //页面标题
			Revisions []struct { //修订者信息
				Revid     int       `json:"revid"`     //修订ID
				Parentid  int       `json:"parentid"`  //父ID
				User      string    `json:"user"`      //修订者
				TimeStamp time.Time `json:"timestamp"` //修订时间
				Comment   string    `json:"comment"`   //修订说明
			} `json:"revisions"`
		} `json:"pages"`
	} `json:"query"`
}

/**
 * @description: 查询页面修订信息
 * @param {string} WikiLink Wiki链接
 * @param {string} title 需要查询的页面标题
 * @return {QueryRevisionsJson}
 * @return {error}
 */
func QueryRevisions(WikiLink string, title string) (QueryRevisionsJson, error) {
	title = url.QueryEscape(title)
	url := WikiLink + "/api.php?action=query&prop=revisions&format=json&titles=" + title
	body, err := Utils.HttpRequest(url)

	var info QueryRevisionsJson
	json.Unmarshal([]byte(body), &info)
	return info, err
}

type QuerySiteinfoGeneralJson struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         struct {
		General struct {
			Mainpage             string     `json:"mainpage"`   //MediaWiki版本
			Base                 string     `json:"base"`       //首页地址
			Sitename             string     `json:"sitename"`   //站点名字
			Logo                 string     `json:"logo"`       //站点LogoURL
			Generator            string     `json:"generator"`  //MediaWiki版本
			Phpversion           string     `json:"phpversion"` //PHP版本
			Phpsapi              string     `json:"phpsapi"`    //PHP运行方式
			Dbtype               string     `json:"dbtype"`     //数据库类型
			Dbversion            string     `json:"dbversion"`  //数据库版本
			Langconversion       string     `json:"langconversion"`
			Titleconversion      string     `json:"titleconversion"`
			Linkprefixcharset    string     `json:"linkprefixcharset"`
			Linkprefix           string     `json:"linkprefix"`
			Linktrail            string     `json:"linktrail"`
			Legaltitlechars      string     `json:"legaltitlechars"`
			Invalidusernamechars string     `json:"invalidusernamechars"`
			Fixarabicunicode     string     `json:"fixarabicunicode"`
			Fixmalayalamunicode  string     `json:"fixmalayalamunicode"`
			GitHash              string     `json:"git-hash"`
			GitBranch            string     `json:"git-branch"`
			Case                 string     `json:"case"`
			Lang                 string     `json:"lang"` //语言
			Fallback             []struct { //其他语言支持
				Code string `json:"code"` //语言代号
			} `json:"fallback"`
			Fallback8BitEncoding string    `json:"fallback8bitEncoding"`
			Writeapi             string    `json:"writeapi"`
			Maxarticlesize       int       `json:"maxarticlesize"`
			Timezone             string    `json:"timezone"`
			Timeoffset           int       `json:"timeoffset"`
			Articlepath          string    `json:"articlepath"`
			Scriptpath           string    `json:"scriptpath"`
			Script               string    `json:"script"`
			Variantarticlepath   bool      `json:"variantarticlepath"`
			Server               string    `json:"server"`     //服务地址
			Servername           string    `json:"servername"` //域名
			Wikiid               string    `json:"wikiid"`
			Time                 time.Time `json:"time"` //Wiki服务器当前时间
			Uploadsenabled       string    `json:"uploadsenabled"`
			Maxuploadsize        int       `json:"maxuploadsize"`
			Minuploadchunksize   int       `json:"minuploadchunksize"`
			Galleryoptions       struct {
				ImagesPerRow   int    `json:"imagesPerRow"`
				ImageWidth     int    `json:"imageWidth"`
				ImageHeight    int    `json:"imageHeight"`
				CaptionLength  string `json:"captionLength"`
				ShowBytes      string `json:"showBytes"`
				ShowDimensions string `json:"showDimensions"`
				Mode           string `json:"mode"`
			} `json:"galleryoptions"`
			Thumblimits []int `json:"thumblimits"`
			Imagelimits []struct {
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"imagelimits"`
			Favicon                     string        `json:"favicon"` //站点favicon
			Centralidlookupprovider     string        `json:"centralidlookupprovider"`
			Allcentralidlookupproviders []string      `json:"allcentralidlookupproviders"`
			Interwikimagic              string        `json:"interwikimagic"`
			Magiclinks                  []interface{} `json:"magiclinks"`
			Categorycollation           string        `json:"categorycollation"`
			Citeresponsivereferences    string        `json:"citeresponsivereferences"`
		} `json:"general"`
	} `json:"query"`
}

/**
 * @description: 查询网站的全部系统信息
 * @param {string} WikiLink Wiki链接
 * @return {QuerySiteinfoGeneralJson}
 * @return {error}
 */
func QuerySiteinfoGeneral(WikiLink string) (QuerySiteinfoGeneralJson, error) {
	url := WikiLink + "/api.php?action=query&meta=siteinfo&siprop=general&format=json"
	body, err := Utils.HttpRequest(url)

	var info QuerySiteinfoGeneralJson
	json.Unmarshal([]byte(body), &info)
	return info, err
}
