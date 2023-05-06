package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func main() {
	// 命令行参数获取下载链接和线程数量
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run downloader.go <url> <threads>")
		return
	}
	url := os.Args[1]
	threads, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid thread count:", os.Args[2])
		return
	}

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
	fmt.Println("File size:", size)

	// 创建文件并分配文件块
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
	for i := 0; i < threads; i++ {
		go func(b [2]int) {
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

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response:", err)
				wg.Done()
				return
			}

			f.Seek(int64(b[0]), 0)
			f.Write(data)

			fmt.Printf("Downloaded %v bytes from %v-%v\n", len(data), b[0], b[1]-1)
			wg.Done()
		}(blocks[i])
	}

	wg.Wait()
	fmt.Println("Download complete.")
}
