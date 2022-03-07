/*
 * @Author: NyanCatda
 * @Date: 2021-10-03 16:45:15
 * @LastEditTime: 2022-01-24 19:45:03
 * @LastEditors: NyanCatda
 * @Description: Post请求封装
 * @FilePath: \ShionBot\src\Utils\HttpRequest\PostRequestJosn.go
 */
package HttpRequest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
 * @description: POST请求封装，传递Json
 * @param {string} url 请求地址
 * @param {string} requestBody 请求内容(Json)
 * @param {[]string} Header 请求头
 * @return {[]byte} 返回内容
 * @return {*http.Response} 请求响应信息
 * @return {error} Error
 */
func PostRequestJson(url string, requestBody string, Header []string) ([]byte, *http.Response, error) {
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range Header {
		Headervalue := strings.Split(value, ":")
		req.Header.Set(Headervalue[0], Headervalue[1])
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		var errbody []byte
		return errbody, resp, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, resp, err
}
