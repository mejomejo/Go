package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sync"
)

var mainch = make(chan string)
var wg sync.WaitGroup

var stop = make(chan bool)

func //main() {

	var URL string
	wg.Add(10)
	for i := 0; i < 250; i += 25 {
		URL = fmt.Sprintf("http://douban.com/movie/top250?start=%d", i)
		go output(URL)
	}

	go func() {
		for i := range mainch {
			fmt.Println(i)
		}
		stop <- true
	}()

	wg.Wait()
	close(mainch)
	<-stop
	fmt.Println("Done!")
}

func fetch(URL string) string {
	reqList, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	reqList.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	//reqList.Header.Add("Cookie", "bid=QZxxL19DKZc; ll=\"118371\"; _vwo_uuid_v2=D5C96C240AF5DF538181AB94F7EF81871|e0919e888c485419fb9d8ffc95ef560b; gr_user_id=70353b69-5a09-4462-85dd-36eb7126e6be; viewed=\"36081529_30280001_27103952_26293546_35781423_6529833_26590745_26912767_26344642_35781427\"; _pk_ref.100001.4cf6=[\"\",\"\",1682590135,\"https://www.google.com.hk/\"]; _pk_id.100001.4cf6=9ead96e17d3b0557.1680014274..1682590135.undefined.; _pk_ses.100001.4cf6=*; ap_v=0,6.0; __utma=30149280.538171965.1680014275.1682568754.1682590135.10; __utmb=30149280.0.10.1682590135; __utmc=30149280; __utmz=30149280.1682590135.10.9.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not provided); __utma=223695111.1397909160.1680014275.1682568754.1682590135.5; __utmb=223695111.0.10.1682590135; __utmc=223695111; __utmz=223695111.1682590135.5.4.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not provided); ct=y")

	client := &http.Client{}
	resp, err1 := client.Do(reqList)
	if err1 != nil {
		fmt.Println(err1)
		return ""
	}
	defer resp.Body.Close()

	body, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
		return ""
	}
	return string(body)
}

func output(URL string) {

	res := fetch(URL)
	handler1 := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))"`)
	title := handler1.FindAllStringSubmatch(res, -1)

	for _, val := range title {
		mainch <- val[1]
	}

	wg.Done()
}
