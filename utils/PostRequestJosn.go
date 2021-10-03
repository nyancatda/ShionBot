package utils

import (
    "net/http"
    "bytes"
	"io/ioutil"
)

func PostRequestJosn(url string,requestBody string) ([]byte) {
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}