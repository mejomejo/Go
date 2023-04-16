package main

import "fmt"

func addn(n int) func(x int) int {
	fun := func(x int) int {
		return x + n
	}
	return fun
}

func main() {
	add3 := addn(3)
	fmt.Println(add3(100))
	fmt.Println(add3(200))
	fmt.Println(add3(300))
}
