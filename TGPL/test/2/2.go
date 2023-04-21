package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex1 sync.Mutex

func printer(s string) {
	mutex1.Lock()
	for _, val := range s {
		fmt.Printf("%c", val)
		time.Sleep(time.Millisecond * 50)
	}
	mutex1.Unlock()
}

func aaa() {
	printer("Hello")
}

func bbb() {
	printer("World")
}

func main() {
	go aaa()
	go bbb()
	time.Sleep(time.Second)
}
