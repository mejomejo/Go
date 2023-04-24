package main

import "log"

type TreeNode struct {
	elem  int
	left  *TreeNode
	right *TreeNode
}

func main() {

}

// 创造一棵空树，或者将一棵搜索树清空
func MakeEmpty(tree *TreeNode) *TreeNode {
	if tree != nil {
		MakeEmpty(tree.left)
		tree.left = nil
		MakeEmpty(tree.right)
		tree.right = nil
	}

	return nil
}

func FindMin(tree *TreeNode) *TreeNode {
	if tree == nil {
		return nil
	}

	for tree.left != nil {
		tree = tree.left
	}

	return tree
}

func FindMax(tree *TreeNode) *TreeNode {
	if tree == nil {
		return nil
	}

	if tree.right == nil {
		return tree
	}

	return FindMax(tree.right)
}

func Find(elem int, tree *TreeNode) *TreeNode {
	if tree == nil {
		return nil
	}

	if tree.elem == elem {
		return tree
	} else if elem > tree.elem {
		return Find(elem, tree.right)
	} else {
		// 查找节点比当前树节点要小
		return Find(elem, tree.left)
	}
}

func Insert(elem int, tree *TreeNode) *TreeNode {
	if tree == nil {
		tree = &TreeNode{}
		tree.left = nil
		tree.right = nil
		tree.elem = elem
	} else if elem < tree.elem {
		tree.left = Insert(elem, tree.left)
	} else if elem > tree.elem {
		tree.right = Insert(elem, tree.right)
	} else {
		// 该节点已经在这颗树中了，我们什么也不做
	}

	return tree
}

func Delete(elem int, tree *TreeNode) *TreeNode {
	if tree == nil {
		log.Fatalf("Cannot find element %v in tree %v\n", elem, tree)
	}

	if elem < tree.elem {
		tree.left = Delete(elem, tree.left)
	} else if elem > tree.elem {
		tree.right = Delete(elem, tree.right)
	} else {
		// 被删除的就是当前节点

		if tree.left != nil && tree.right != nil {
			// 被删除的节点有两个子节点
			tmpNode := FindMin(tree.right)
			tree.elem = tmpNode.elem
			tree.right = Delete(tmpNode.elem, tree.right)
		} else {
			// 被删除的节点有0个或者1个子节点
			if tree.left == nil {
				tree = tree.right
			} else if tree.right == nil {
				tree = tree.left
			} else {
				tree = nil
			}
		}
	}

	return tree
}
