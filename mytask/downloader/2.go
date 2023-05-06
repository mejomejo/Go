package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var message = make(chan string, 5)
var stop = make(chan bool)

func main() {

	url := "https://speed.hetzner.de/100MB.bin"

	//QQ音乐："https://dldir1.qq.com/music/clntupate/QQMusic_YQQDLPage.exe"

	// 获取文件大小
	resp, err := http.Head(url)
	if err != nil {
		fmt.Println("Error getting file size:", err)
		return
	}
	defer resp.Body.Close()

	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		fmt.Println("Error parsing Content-Length header:", err)
		return
	}
	var mbsize float64 = float64(size) / 1024 / 1024
	fmt.Printf("File size:%.2fMB\n\n", mbsize)

	threads := 10

	f, err := os.Create("downloaded_file")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	blockSize := size / threads
	blocks := make([][2]int, threads)
	for i := range blocks {
		blocks[i][0] = i * blockSize
		blocks[i][1] = (i + 1) * blockSize
		if i == threads-1 {
			blocks[i][1] = size
		}
	}

	// 开始下载
	var wg sync.WaitGroup
	wg.Add(threads)
	go func() {
		for i := range message {
			fmt.Println(i)
		}

		fmt.Println("Download complete.")
		stop <- true

	}()
	for i := 0; i < threads; i++ {
		go func(b [2]int, i int) {

			message <- fmt.Sprintf("I'm %d,mytask is to download  %v-%v", i, b[0]/1024/1024, (b[1]-1)/1024/1024)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				wg.Done()
				return
			}
			req.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", b[0], b[1]-1))

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Println("Error downloading file:", err)
				wg.Done()
				return
			}
			defer resp.Body.Close()

			data, err := io.ReadAll(resp.Body)
			if err != nil && len(data) != 0 {
				fmt.Println("Error reading response:", err)
				wg.Done()
				return
			}

			f.Seek(int64(b[0]), 0)
			f.Write(data)

			message <- fmt.Sprintf("Downloaded %v MB from %v-%v", len(data)/1024/1024, b[0]/1024/1024, (b[1]-1)/1024/1024)
			wg.Done()
		}(blocks[i], i)
	}

	wg.Wait()
	close(message)
	<-stop
}
