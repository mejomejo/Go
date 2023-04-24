package main

import "fmt"

func QuickSort(a []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	t := a[left]
	for i < j {
		for a[j] >= t && i < j {
			j--
		}
		for a[i] <= t && i < j {
			i++
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	a[left] = a[j]
	a[j] = t
	QuickSort(a, left, i-1)
	QuickSort(a, i+1, right)
}

func main() {
	a := []int{11, 97, 33, 84, 42, 23, 42, 25, 31, 13, 44, 32, 18, 50, 27, 84, 20, 83, 73}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
