package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	for _, i := range os.Args[1:] {
		s += i + " "
	}
	fmt.Println(s)
}
