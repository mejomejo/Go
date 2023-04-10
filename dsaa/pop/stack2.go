package main

import "fmt"

type stack struct {
	val  int
	next *stack
}

func create() *stack {
	t := new(stack)
	t.val = 0
	t.next = nil
	return t
}

func push(sentinal *stack, val int) {
	t := &stack{val: val, next: sentinal.next}
	sentinal.next = t
	sentinal.val++
}

func show(sentinal *stack) {
	for t := sentinal.next; t != nil; t = t.next {
		fmt.Printf("%d --> ", t.val)
	}
	fmt.Println("nil")
}

func pop(sentinal *stack) int {
	a := sentinal.next.val
	sentinal.next = sentinal.next.next
	sentinal.val--
	return a
}

func size(sentinal *stack) int {
	return sentinal.val
}

func isEmpty(sentinal *stack) bool {
	return sentinal.val == 0
}
func destroy(sentinal *stack) {
	sentinal = nil
}

func main() {
	sentinal := create()
	push(sentinal, 10)
	push(sentinal, 20)
	push(sentinal, 30)
	push(sentinal, 40)
	show(sentinal)
	fmt.Println(size(sentinal))
	pop(sentinal)
	show(sentinal)
	fmt.Println(sentinal.val)
	destroy(sentinal)
}
