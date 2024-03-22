package binarytree

import (
	"fmt"

	"github.com/basokant/go-data-structures/tree"
)

func (bt BinaryTree[T]) Search(data T) (tree.TreeNode[T], error) {
	return bt.SearchWithOrder(data, Preorder)
}

func (bt BinaryTree[T]) SearchWithOrder(data T, order Order) (tree.TreeNode[T], error) {
	var foundNode tree.TreeNode[T]
	var empty *BinaryTreeNode[T]

	if bt.root == empty {
		bt.root = &BinaryTreeNode[T]{data: data, children: []*BinaryTreeNode[T]{}}
	}

	bt.Traverse(func(node tree.TreeNode[T]) bool {
		left, right := Left(node), Right(node)

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
