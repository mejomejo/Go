package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 归并排序
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 将数组分成两部分
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// 合并两个有序数组
	return merge(left, right)
}

// 合并两个有序数组
func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}

func GenerateRandomArray(n int) []int {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100000) + 1 // 生成 0-99 之间的随机整数
	}
	return arr
}

func main() {
	arr := GenerateRandomArray(10000)
	sortedArr := MergeSort(arr)
	fmt.Println(sortedArr)
}
