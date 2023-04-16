package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("comma(\"1234567890\"): %v\n", comma("1234567890"))
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s) % 3
	if n == 0 {
		n = 3
	}

	buf.WriteString(s[:n])

	for n < len(s) {
		buf.WriteByte(',')
		buf.WriteString(s[n : n+3])
		n += 3
	}

	return buf.String()
}
