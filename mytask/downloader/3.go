package main

import (
	"fmt"
)

var ch = make(chan string)

func main() {

	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("Hello %d", i)
		}
		close(ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}

}
