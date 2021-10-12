package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"log"
)

func HttpRequest(url string) ([]byte,error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		log.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		log.Println(err)
	}
	return body,err
}