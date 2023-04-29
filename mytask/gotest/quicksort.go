package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left, right := make([]int, 0, len(arr)), make([]int, 0, len(arr))
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}
	left, right = quicksort(left), quicksort(right)
	return append(append(left, pivot), right...)
}

func main() {
	arr := []int{5, 2, 6, 3, 1, 4}
	sorted := quicksort(arr)
	fmt.Println(sorted)
}
