package binarytree

import (
	"errors"

	"github.com/basokant/go-data-structures/tree"
	"golang.org/x/exp/constraints"
)

type BinaryTree[T constraints.Ordered] struct {
	root *BinaryTreeNode[T]
}

func NewBinaryTree[T constraints.Ordered]() TraversableTree[T] {
	return &BinaryTree[T]{}
}

func (bt BinaryTree[T]) Root() (tree.TreeNode[T], error) {
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

func Height[T constraints.Ordered](node tree.TreeNode[T]) int {
	if node == nil {
		return 0
	}

	leftHeight := Height(Left(node))
	rightHeight := Height(Right(node))
	return max(leftHeight, rightHeight) + 1
}

func (bt BinaryTree[T]) Len() int {
	len := 0

	bt.Traverse(func(_ tree.TreeNode[T]) bool {
		len++
		return true
	}, Inorder)

	return len
}
