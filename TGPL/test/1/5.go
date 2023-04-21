package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fp, err0 := os.OpenFile("1.go", os.O_RDONLY, 6)
	if err0 != nil {
		fmt.Println(err0)
		return
	}
	reader := bufio.NewReader(fp)
	count := 0
	for {
		buf, err1 := reader.ReadString('\n')
		if err1 != nil {
			//fmt.Println(err1)
			break
		}
		line := strings.Fields(buf)
		for _, val := range line {
			if strings.Contains(val, "func") {
				count++
			}
		}
	}
	fmt.Println(count)
}
