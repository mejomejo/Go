package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type Spider struct {
	url   string
	Maxgo int
	Depth int
}

var mainch = make(chan int)

var ch = make(chan int, 8)

func main() {

	// go func() {
	// 	ch <- 1
	// }()

	aaa := &Spider{
		url:   "https://ssr1.scrape.center/page/",
		Maxgo: 5,
		Depth: 1,
	}

	for i := 1; i <= 11; i++ {
		if aaa.Maxgo > 0 {
			go get(aaa, i)
			aaa.Maxgo--
		} else {
			get(aaa, i)
		}
	}
	for i := 0; i < 11; i++ {
		<-mainch
	}
	// <-ch
}

func get(aaa *Spider, i int) {

	// for {
	// 	t := <-ch
	// 	if t == i {
	// 		break
	// 	} else {
	// 		ch <- t
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }

	url := fmt.Sprintf(aaa.url+"%d", i)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	aaa.Maxgo++

	content, err := io.ReadAll(resp.Body)

	handler1 := regexp.MustCompile(`"m-b-sm">(.*?)</h2>`)
	titles := handler1.FindAllStringSubmatch(string(content), -1)

	handler2 := regexp.MustCompile(`(.*?)</p>`)
	where := handler2.FindAllStringSubmatch(string(content), -1)

	for index, title := range titles {
		fmt.Println(strings.Split(title[1], "-")[0], where[makeindex(index)][1])
	}
	// go func() {
	// 	ch <- i + 1
	// }()
	go func() {
		mainch <- 0
	}()
}

func makeindex(i int) int {
	if i%2 == 0 {
		return i
	}
	return i + 1
}
