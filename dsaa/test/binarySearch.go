package main

import "fmt"

func Find(a []int, val int) int {
	x, y := 0, len(a)-1
	for x < y {
		mid := x + (y-x)/2
		if a[mid] == val {
			return mid
		}
		if a[mid] > val {
			y = mid - 1
		}
		if a[mid] < val {
			x = mid + 1
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := Find(a, 2)
	fmt.Println(b)
}
