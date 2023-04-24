package main

import "testing"

func TestInsert(t *testing.T) {
	var tree *TreeNode = nil

	tree = Insert(3, tree)
	tree = Insert(4, tree)
	tree = Insert(1, tree)
	tree = Insert(2, tree)

	// 构造出来的树
	//    3
	//   / \
	//   1  4
	//    \
	//     2

	assert.Equal(t, 4, FindMax(tree).elem)
	assert.Equal(t, 1, FindMin(tree).elem)

	assert.Nil(t, Find(4, tree).left)
	assert.Nil(t, Find(4, tree).right)
}

func TestDeleteRight(t *testing.T) {
	var tree *TreeNode = nil

	// 构造出来的树
	//     6
	//    / \
	//   2  8
	//  /\
	// 1 4
	//   \
	//    5
	tree = Insert(6, tree)
	tree = Insert(2, tree)
	tree = Insert(8, tree)
	tree = Insert(1, tree)
	tree = Insert(4, tree)
	tree = Insert(5, tree)

	// 删除了节点4之后的树
	//     6
	//    / \
	//   2  8
	//  /\
	// 1 5
	tree = Delete(4, tree)

	assert.Equal(t, 5, Find(2, tree).right.elem)
}

func TestDeleteTwoChild(t *testing.T) {
	var tree *TreeNode = nil

	// 删除之前的树
	//     6
	//    / \
	//   2  8
	//  /\
	// 1 5
	//  /
	// 3
	// \
	//  4
	tree = Insert(6, tree)
	tree = Insert(2, tree)
	tree = Insert(8, tree)
	tree = Insert(1, tree)
	tree = Insert(5, tree)
	tree = Insert(3, tree)
	tree = Insert(4, tree)

	// 删除节点2之后的树
	//     6
	//    / \
	//   3  8
	//  /\
	// 1 5
	//  /
	// 4
	tree = Delete(2, tree)

	node := Find(3, tree)
	assert.Equal(t, 1, node.left.elem)
	assert.Equal(t, 5, node.right.elem)
}
