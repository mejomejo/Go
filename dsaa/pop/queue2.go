package main

import "fmt"

type Queue struct {
	element [8]int
	head    int
	tail    int
}

func create() *Queue {
	t := new(Queue)
	t.head = 0
	t.tail = 0
	return t
}

func isEmpty(q *Queue) bool {
	return q.head == q.tail
}

func isFull(q *Queue) bool {
	return (q.tail+1)%8 == q.head
}

func add(q *Queue, val int) int {
	if isFull(q) {
		return -1
	}
	q.element[q.tail] = val
	q.tail = (q.tail + 1) % 8
	return 0
}

func sub(q *Queue) int {
	if isEmpty(q) {
		return -1
	}
	a := q.element[q.head]
	q.head = (q.head + 1) % 8
	return a
}

func show(q *Queue) {
	for i := q.head; i != q.tail; i = (i + 1) % 8 {
		fmt.Printf("%d --> ", q.element[i])
	}
	fmt.Println("nil")
}

func main() {
	q := create()
	add(q, 1)
	add(q, 3)
	add(q, 5)
	add(q, 7)
	add(q, 9)
	add(q, 2)
	fmt.Printf("add(q, 4): %v\n", add(q, 4))
	fmt.Printf("add(q, 6): %v\n", add(q, 6))
	show(q)
	sub(q)
	sub(q)
	sub(q)
	sub(q)
	show(q)
	sub(q)
	sub(q)
	show(q)
	fmt.Printf("sub(q): %v\n", sub(q))
	fmt.Printf("sub(q): %v\n", sub(q))
}
