package main

import "fmt"

type MyLinkedList[T any] struct {
	sentinal Node[T]
}

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func (this *MyLinkedList[T]) Add(Val T) {
	t := Node[T]{
		Val:  Val,
		Next: this.sentinal.Next,
	}
	this.sentinal.Next = &t
}

func (this *MyLinkedList[T]) Del(Val T) T {
	if this.isEmpty() {

	}
}

func (this *MyLinkedList[T]) Isempty() bool {
	return this.sentinal.Next == nil
}

func (this *MyLinkedList[T]) Show() {
	t := this.sentinal.Next
	for t != nil {
		fmt.Print(t.Val, "->")
		t = t.Next
	}

	fmt.Println("nil")
}

func main() {
	strList := new(MyLinkedList[string])
	strList.Add("lihua")
	strList.Add("woshi ")
	strList.Add("nihao ")
	strList.Show()

}
