package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	URL := "http://douban.com/movie/top250"
	reqList, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	reqList.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	reqList.Header.Add("Cookie", "bid=QZxxL19DKZc; ll=\"118371\"; _vwo_uuid_v2=D5C96C240AF5DF538181AB94F7EF81871|e0919e888c485419fb9d8ffc95ef560b; gr_user_id=70353b69-5a09-4462-85dd-36eb7126e6be; viewed=\"36081529_30280001_27103952_26293546_35781423_6529833_26590745_26912767_26344642_35781427\"; _pk_ref.100001.4cf6=[\"\",\"\",1682590135,\"https://www.google.com.hk/\"]; _pk_id.100001.4cf6=9ead96e17d3b0557.1680014274..1682590135.undefined.; _pk_ses.100001.4cf6=*; ap_v=0,6.0; __utma=30149280.538171965.1680014275.1682568754.1682590135.10; __utmb=30149280.0.10.1682590135; __utmc=30149280; __utmz=30149280.1682590135.10.9.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not provided); __utma=223695111.1397909160.1680014275.1682568754.1682590135.5; __utmb=223695111.0.10.1682590135; __utmc=223695111; __utmz=223695111.1682590135.5.4.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not provided); ct=y")

	client := &http.Client{}
	resp, err1 := client.Do(reqList)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(string(body))
}
