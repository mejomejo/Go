package main

import "fmt"

type stack struct {
	element []int
	top     int
}

func create() *stack {
	t := new(stack)
	t.top = -1
	t.element = make([]int, 8)
	return t
}

func isEmpty(m *stack) bool {
	return m.top == -1
}
func push(m *stack, val int) {
	if m.top == 7 {
		return
	}
	m.top++
	m.element[m.top] = val
}
func pop(m *stack) int {
	if m.top == -1 {
		return -1
	}
	m.top--
	return m.element[m.top+1]
}

func size(m *stack) int {
	return m.top + 1
}

func destroy(m *stack) {
	m = nil
}

func main() {
	m := create()
	fmt.Println(isEmpty(m))
	fmt.Println(size(m))
	push(m, 3)
	push(m, 6)
	fmt.Println("size:", size(m))
	fmt.Println(pop(m))
	fmt.Println("size:", size(m))
	fmt.Println(pop(m))
	fmt.Println(pop(m)) // -1 means pop error
	destroy(m)
}
