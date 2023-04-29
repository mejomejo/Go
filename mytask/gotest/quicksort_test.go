package main

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	arr := []int{5, 2, 6, 3, 1, 4}
	sorted := quicksort(arr)
	expected := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("quicksort(%v) = %v, want %v", arr, sorted, expected)
	}
}
