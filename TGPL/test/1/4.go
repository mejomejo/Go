package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fp, err0 := os.OpenFile("1.go", os.O_RDWR, 6)
	if err0 != nil {
		fmt.Println(err0)
		return
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	for {
		buf, err1 := reader.ReadBytes('\n')
		if err1 != nil {
			fmt.Println(err1)
			break
		}
		fmt.Print(string(buf))
	}
}
