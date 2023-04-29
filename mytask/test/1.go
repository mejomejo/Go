package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var mainch = make(chan int)

var ch = make(chan int, 1)

func main() {

	// go func() {
	// 	ch <- 1
	// }()

	for i := 1; i <= 11; i++ {
		go get(i)
	}
	for i := 0; i < 11; i++ {
		<-mainch
	}
	// <-ch
}

func get(i int) {

	// for {
	// 	t := <-ch
	// 	if t == i {
	// 		break
	// 	} else {
	// 		ch <- t
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }

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

	handler2 := regexp.MustCompile(`(.*?)</p>`)
	where := handler2.FindAllStringSubmatch(string(content), -1)

	for index, title := range titles {
		fmt.Println(strings.Split(title[1], "-")[0], where[makeindex(index)][1])
	}
	//ch <- i + 1
	mainch <- 0
}

func makeindex(i int) int {
	if i == 0 {
		return i
	}
	return i + 1
}
