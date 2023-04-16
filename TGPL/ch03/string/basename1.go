package main

import "fmt"

func main() {
	fmt.Printf("basename(\"/home/dyx/go/gopath/main.go\"): %v\n", basename("/home/dyx/go/gopath/main.go"))
}
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
