package binarytree

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type BinaryTree[T constraints.Ordered] struct {
	root *BinaryTreeNode[T]
}

func NewBinaryTree[T constraints.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

func (bt BinaryTree[T]) Root() (*BinaryTreeNode[T], error) {
	if bt.root == nil {
		return nil, errors.New("binary tree does not have root")
	}
	return bt.root, nil
}

func (bt *BinaryTree[T]) SetRoot(node *BinaryTreeNode[T]) {
	bt.root = node
}

func (bt *BinaryTree[T]) Height() int {
	return Height(bt.root)
}

func Height[T constraints.Ordered](node *BinaryTreeNode[T]) int {
	if node == nil {
		return 0
	}

	leftHeight := Height(node.Left)
	rightHeight := Height(node.Right)
	return max(leftHeight, rightHeight) + 1
}

func (bt BinaryTree[T]) Len() int {
	len := 0

	bt.Traverse(func(_ *BinaryTreeNode[T]) bool {
		len++
		return true
	}, Inorder)

	return len
}
