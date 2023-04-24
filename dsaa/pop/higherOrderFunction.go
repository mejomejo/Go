package main

import (
	"fmt"
	"math"
)

func addn(n int) func(x int) int {
	fun := func(x int) int {
		return x + n
	}
	return fun
}

func main() {
	square := make1(3)
	fmt.Println(square(4))
}

func make1(n float64) func(x float64) float64 {
	fun := func(x float64) float64 {
		return math.Pow(x, n)
	}
	return fun
}
