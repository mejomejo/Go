package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func main() {
	go func() {
		ch <- 1
	}()

	for i := 0; i < 10; i++ {
		go test(i + 1)
	}

	time.Sleep(1 * time.Second)
}

func test(i int) {

	for {
		t := <-ch
		if t == i {
			break
		} else {
			ch <- t
		}
	}
	fmt.Println(i)

	ch <- i + 1
}
