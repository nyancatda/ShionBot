/*
 * @Author: NyanCatda
 * @Date: 2021-10-03 16:40:29
 * @LastEditTime: 2022-01-24 19:44:21
 * @LastEditors: NyanCatda
 * @Description: Get请求封装
 * @FilePath: \ShionBot\src\Utils\HttpRequest\GetRequest.go
 */
package HttpRequest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
 * @description: GET请求封装
 * @param {string} url 请求地址
 * @param {[]string} Header 请求头
 * @return {[]byte} 返回内容
 * @return {*http.Response} 请求响应信息
 * @return {error} Error
 */
 func GetRequest(url string, Header []string) ([]byte, *http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range Header {
		Headervalue := strings.Split(value, ":")
		req.Header.Set(Headervalue[0], Headervalue[1])
	}

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
