package main

import (
	"fmt"
)

func print() {
	fmt.Printf("%d go\n", 1)
	//time.Sleep(1 * time.Millisecond)
}

func main() {
	go print()
	//fmt.Println("1 ")

	//time.Sleep(1 * time.Second)
}
