package main

import (
	"fmt"
	"strings"
)

func count(a string) map[string]int {
	map1 := make(map[string]int)
	buf := strings.Fields(a)
	for _, val := range buf {
		map1[val]++
	}
	return map1
}

func main() {
	a := count("Hello World I want to tell  you something you like very much it's Hello World")
	fmt.Println(a)
}
