package main

import "fmt"

func getNext(pattern string) []int {
	m := len(pattern)
	next := make([]int, m)
	j := 0
	for i := 1; i < m; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = next[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		next[i] = j
	}
	return next
}

func kmp(text, pattern string) int {
	n := len(text)
	m := len(pattern)
	if m == 0 {
		return 0
	}
	next := getNext(pattern)
	j := 0
	for i := 0; i < n; i++ {
		for j > 0 && text[i] != pattern[j] {
			j = next[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func main() {
	text := "abacadeabcfejjiiu"
	pattern := "abac"
	fmt.Println(getNext(pattern))
	pos := kmp(text, pattern)
	if pos == -1 {
		fmt.Println("Pattern not found.")
	} else {
		fmt.Printf("Pattern found at position %d.\n", pos)
	}
}
