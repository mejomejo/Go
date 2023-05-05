package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var ARRLEN = 10000000
var EXTENT = 10000

func main() {
	a := randarr()
	sort.Ints(a)
	fmt.Println(a)
}

func randarr() []int {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, ARRLEN)
	for i := 0; i < ARRLEN; i++ {
		arr[i] = rand.Intn(EXTENT) + 1
	}
	return arr
}

func bubble(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func quickSort(a []int) {
	if len(a) < 2 {
		return
	}
	left, right := 0, len(a)-1
	t := a[0]
	for left < right {
		for left < right && a[right] >= t {
			right--
		}
		for left < right && a[left] <= t {
			left++
		}
		if left < right {
			a[left], a[right] = a[right], a[left]
		}
	}
	a[0] = a[left]
	a[left] = t
	quickSort(a[:left])
	quickSort(a[left+1:])
}
