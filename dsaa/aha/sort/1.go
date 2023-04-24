package main

import "fmt"

func main() {
	a := [11]int{}
	for i := 0; i < 5; i++ {
		var t int
		fmt.Scanf("%d", &t)
		a[t]++
	}
	for i := 0; i <= 10; i++ {
		for j := 1; j <= a[i]; j++ {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
