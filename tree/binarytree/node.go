package binarytree

import (
	"golang.org/x/exp/constraints"
)

type BinaryTreeNode[T constraints.Ordered] struct {
	data  T
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}

func NewBinaryTreeNode[T constraints.Ordered](data T, left *BinaryTreeNode[T], right *BinaryTreeNode[T]) *BinaryTreeNode[T] {
	return &BinaryTreeNode[T]{data, left, right}
}

func (node BinaryTreeNode[T]) Data() T {
	return node.data
}

func (node *BinaryTreeNode[T]) SetData(data T) {
	node.data = data
}

func (node BinaryTreeNode[T]) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}
