package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostRequestJosn(url string,requestBody string) ([]byte,*http.Response,error) {
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		var errbody []byte
		return errbody,resp,err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body,resp,err
}