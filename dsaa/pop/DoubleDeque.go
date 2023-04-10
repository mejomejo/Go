package main

import "fmt"

type Node struct {
	prior *Node
	next  *Node
	val   int
}

func create() *Node {
	t := new(Node)
	t.next = t
	t.prior = t
	t.val = 0 //size
	return t
}

func addFirst(sentinal *Node, val int) {
	t := new(Node)
	sentinal.val++
	t.val = val
	t.next = sentinal.next
	sentinal.next.prior = t
	t.prior = sentinal
	sentinal.next = t
}

func addLast(sentinal *Node, val int) {
	t := new(Node)
	sentinal.val++
	t.val = val
	t.prior = sentinal.prior
	sentinal.prior.next = t
	sentinal.prior = t
	t.next = sentinal
}

func removeFirst(sentinal *Node) {
	sentinal.val--
	sentinal.next.next.prior = sentinal
	sentinal.next = sentinal.next.next
}

func removeLast(sentinal *Node) {
	sentinal.val--
	sentinal.prior.prior.next = sentinal
	sentinal.prior = sentinal.prior.prior
}

func size(sentinal *Node) int {
	return sentinal.val
}

func get(sentinal *Node, i int) int {
	if i > sentinal.val {
		return -1
	}
	t := sentinal
	for j := i; j > 0; j-- {
		t = t.next
	}
	return t.val
}

func show(sentinal *Node) {
	for t := sentinal.next; t != sentinal; t = t.next {
		fmt.Printf("%d -> ", t.val)
	}
	fmt.Println("sentinal")
}

func isEmpty(sentinal *Node) bool {
	return sentinal.val == 0
}

func destroy(sentinal **Node) {
	*sentinal = nil
}

func main() {
	sentinal := create()
	fmt.Println(isEmpty(sentinal))
	addFirst(sentinal, 10)
	addFirst(sentinal, 20)
	addFirst(sentinal, 30)
	show(sentinal)
	addLast(sentinal, 10)
	addLast(sentinal, 20)
	addLast(sentinal, 30)
	show(sentinal)
	fmt.Println(sentinal.val)
	fmt.Println(get(sentinal, 2))
	removeFirst(sentinal)
	show(sentinal)
	fmt.Println(get(sentinal, 2))
	fmt.Println(sentinal.val)
	removeLast(sentinal)
	show(sentinal)
	fmt.Printf("%v\n", *sentinal)
	fmt.Println(isEmpty(sentinal))
	destroy(&sentinal)
	//isEmpty(sentinal)
}
