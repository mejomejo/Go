package main

import (
	"fmt"
	"os"
)

func main() {
	for i, j := range os.Args[1:] {
		fmt.Printf("index: %d, args: %s\n", i, j)
	}
}
