package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var ARRLEN = 100000
var EXTENT = 10000
var MAXGO = 40
var ch = make(chan []int)
var returnch = make(chan int)
var res [][]int

func main() {
	myarr := randarr()
	test := mysplit(myarr)
	go merge(test)
	for _, val := range test {
		go myquick(val, true)
	}

	<-returnch
	// sort.Ints(myarr)
	// fmt.Println(myarr)
}

func randarr() []int {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, ARRLEN)
	for i := 0; i < ARRLEN; i++ {
		arr[i] = rand.Intn(EXTENT) + 1
	}
	return arr
}

func mysplit(myarr []int) [][]int {
	ll := len(myarr) / MAXGO
	var test [][]int
	for i := 0; i < len(myarr); i += ll {
		end := i + ll
		if end > len(myarr) {
			end = len(myarr)
		}
		test = append(test, myarr[i:end])
	}
	return test
}

func myquick(a []int, root bool) {
	if root {
		defer myreturn(a)
	}

	if len(a) < 2 {
		return
	}
	left, right := 0, len(a)-1
	t := a[left]
	for left < right {
		for left < right && a[right] >= t {
			right--
		}
		if left < right {
			a[left] = a[right]
		}
		for left < right && a[left] <= t {
			left++
		}

		if left < right {
			a[right] = a[left]
		}
	}
	a[left] = t
	myquick(a[:left], false)
	myquick(a[left+1:], false)
}

func myreturn(a []int) {
	ch <- a
}

func merge(test [][]int) {
	for i := 0; i < MAXGO; i++ {
		<-ch
	}

	res := make([]int, len(test)*len(test[0]))
	i := 0
	for _, val := range test {
		copy(res[i:], val)
		i += len(val)
	}
	sort.Ints(res)
	fmt.Println(res)

	returnch <- 1
}
