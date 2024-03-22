package binarytree

import (
	"fmt"
)

func (bt BinaryTree[T]) Search(data T) (*BinaryTreeNode[T], error) {
	return bt.SearchWithOrder(data, Preorder)
}

func (bt BinaryTree[T]) SearchWithOrder(data T, order Order) (*BinaryTreeNode[T], error) {
	var foundNode *BinaryTreeNode[T]
	var empty *BinaryTreeNode[T]

	if bt.root == empty {
		bt.root = &BinaryTreeNode[T]{data: data}
	}

	bt.Traverse(func(node *BinaryTreeNode[T]) bool {
		left, right := node.Left, node.Right

		if left.Data() == data {
			foundNode = left
			return false
		} else if right.Data() == data {
			foundNode = right
			return false
		}

		return true
	}, order)

	if foundNode == empty {
		return nil, fmt.Errorf("")
	}

	return foundNode, nil
}
