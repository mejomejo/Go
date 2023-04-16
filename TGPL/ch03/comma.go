package main

import "fmt"

func main() {
	fmt.Printf("comma(\"1234567890\"): %v\n", comma("1234567890"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
