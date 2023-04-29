package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	for i := 1; i <= 11; i++ {
		get(i)
	}
}

func get(i int) {

	url := fmt.Sprintf("https://ssr1.scrape.center/page/%d", i)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	handler1 := regexp.MustCompile(`"m-b-sm">(.*?)</h2>`)
	titles := handler1.FindAllStringSubmatch(string(content), -1)

	for index, title := range titles {
		fmt.Println((i-1)*10+index+1, strings.Split(title[1], "-")[0])
	}
}
