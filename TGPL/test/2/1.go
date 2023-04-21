package main

import (
	"fmt"
	"time"
)

func producer(a chan<- int) {
	for i := 1; i < 11; i++ {
		a <- i * i
		fmt.Printf("%c", 'p')
	}
	close(a)
}

func consumer(a <-chan int) {
	for val := range a {
		fmt.Println(val)
		fmt.Printf("%c", 'c')
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}
