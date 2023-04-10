package main

import (
	"fmt"
)

type Node struct {
	val   int
	prior *Node
	next  *Node
}

type Deque struct {
	sentinal Node
}

func (this *Deque) init() {
	this.sentinal.next = &this.sentinal
	this.sentinal.prior = &this.sentinal
	this.sentinal.val = 0
}

func (this *Deque) size() int {
	return this.sentinal.val
}

func (this *Deque) addFirst(val int) {
	t := new(Node)
	t.val = val
	t.next = this.sentinal.next
	this.sentinal.next.prior = t
	t.prior = &this.sentinal
	this.sentinal.next = t
	this.sentinal.val++
}

func (this *Deque) addLast(val int) {
	t := new(Node)
	t.val = val
	t.prior = this.sentinal.prior
	this.sentinal.prior.next = t
	t.next = &this.sentinal
	this.sentinal.prior = t
	this.sentinal.val++
}

func (this *Deque) removeFirst() {
	this.sentinal.next.next.prior = &this.sentinal
	this.sentinal.next = this.sentinal.next.next
	this.sentinal.val--
}

func (this *Deque) removeLast() {
	this.sentinal.prior.prior.next = &this.sentinal
	this.sentinal.prior = this.sentinal.prior.prior
	this.sentinal.val--
}

func (this *Deque) destroy() {
	this.sentinal.prior = nil
	this.sentinal.next = nil
	this.sentinal.val = 0
}

func (this *Deque) isEmpty() bool {
	return this.sentinal.val == 0
}

func (this *Deque) show() {
	fmt.Printf("sentinal -> ")
	for t := this.sentinal.next; t != &this.sentinal; t = t.next {
		fmt.Printf("%d -> ", t.val)
	}
	fmt.Println("sentinal")
}

func main() {
	t := new(Deque)
	t.init()
	t.addFirst(100)
	t.addFirst(200)
	t.addFirst(300)
	t.show()
	t.addLast(400)
	t.addLast(500)
	t.addLast(600)
	t.show()
	fmt.Println(t)
	t.destroy()
	fmt.Println(t)

}
