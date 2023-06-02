package main

import (
	"fmt"
)

type BstTree struct {
	Val   int
	Left  *BstTree
	Right *BstTree
}

func create(val int) *BstTree {
	return &BstTree{
		Val: val,
	}
}

func add(val int, root *BstTree) *BstTree {

	if root == nil {
		return &BstTree{
			Val: val,
		}
	}
	if root.Val > val {
		root.Left = add(val, root.Left)
	}

	if root.Val < val {
		root.Right = add(val, root.Right)
	}

	return root
}

func search(root *BstTree, val int) *BstTree {
	if root == nil {
		return nil
	}
	if root.Val > val {
		return search(root.Left, val)
	}

	if root.Val < val {
		return search(root.Right, val)
	}
	return root
}

func printInOrder(root *BstTree) {
	if root == nil {
		return
	}
	printInOrder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInOrder(root.Right)
}

func printPreOrder(root *BstTree) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printPreOrder(root.Left)
	printPreOrder(root.Right)
}

func printReOrder(root *BstTree) {
	if root == nil {
		return
	}

	printReOrder(root.Left)
	printReOrder(root.Right)

	fmt.Printf("%d ", root.Val)
}

func findMax(root *BstTree) *BstTree {
	if root.Right == nil {
		return root
	}
	return findMax(root.Right)
}

func del(root *BstTree, val int) *BstTree {
	if root == nil {
		return root
	}

	if root.Val > val {
		root.Left = del(root.Left, val)
	} else if root.Val < val {
		root.Right = del(root.Right, val)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			a := findMax(root.Left)
			root.Val = a.Val
			root.Left = del(root.Left, a.Val)
		}
	}

	return root

}

func main() {
	tree := create(10)
	add(20, tree)
	add(40, tree)
	add(30, tree)
	add(15, tree)
	add(50, tree)
	add(99, tree)
	add(1, tree)
	printInOrder(tree)
	fmt.Println()

	del(tree, 1)
	tree = del(tree, 10)
	printInOrder(tree)
}
