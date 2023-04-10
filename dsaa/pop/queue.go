package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func create() *Queue {
	t := new(Queue)
	return t
}
func show(q *Queue) {
	for i := q.head; i != nil; i = i.next {
		fmt.Printf("%d --> ", i.val)
	}
	fmt.Println("nil")
}
func add(q *Queue, val int) {
	t := new(Node)
	t.val = val
	t.next = nil
	if q.head == nil {
		q.head = t
		q.tail = t
	} else {
		q.tail.next = t
		q.tail = t
	}
}
func sub(q *Queue) int {
	a := q.head.val
	q.head = q.head.next
	return a
}
func isEmpty(q *Queue) bool {
	return q.head == nil
}
func size(q *Queue) int {
	if isEmpty(q) {
		return 0
	}
	i := 1
	for p := q.head; p != q.tail; p = p.next {
		i++
	}
	return i
}
func main() {
	q := create()
	fmt.Println(isEmpty(q))
	add(q, 100)
	add(q, 200)
	add(q, 300)
	fmt.Println(size(q))
	show(q)
	sub(q)
	show(q)
	fmt.Println(isEmpty(q))
	fmt.Println(q)
}
