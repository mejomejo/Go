package main

import "fmt"

type Status int
type Tree struct {
	val         Status
	left, right *Tree
}

func preOrderTraverse(root *Tree) {
	if root == nil {
		return
	}
	fmt.Printf("%d,", root.val)
	preOrderTraverse(root.left)
	preOrderTraverse(root.right)
}

func inOrderTraverse(root *Tree) {
	if root == nil {
		return
	}
	preOrderTraverse(root.left)
	fmt.Printf("%d,", root.val)
	preOrderTraverse(root.right)
}

func postOrderTraverse(root *Tree) {
	if root == nil {
		return
	}
	preOrderTraverse(root.left)
	preOrderTraverse(root.right)
	fmt.Printf("%d,", root.val)
}

func main() {
	x := &Tree{val: 1}
	x.left = &Tree{val: 2}
	x.right = &Tree{val: 3}
	fmt.Println(x)
	preOrderTraverse(x)
	fmt.Println()
	inOrderTraverse(x)
	fmt.Println()
	postOrderTraverse(x)
}
