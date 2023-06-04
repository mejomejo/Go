package main

import "fmt"

func bf(a, b string) int {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j == len(b) {
		return i - j
	}
	return -1
}

func main() {
	a := "abacadeabcfejjiiu"
	b := "bcfejj"
	c := bf(a, b)
	fmt.Println(c)
}
