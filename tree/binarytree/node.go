package binarytree

import (
	"github.com/basokant/go-data-structures/tree"
	"golang.org/x/exp/constraints"
)

type BinaryTreeNode[T constraints.Ordered] struct {
	data     T
	children []*BinaryTreeNode[T]
}

func NewBinaryTreeNode[T constraints.Ordered](data T, left *BinaryTreeNode[T], right *BinaryTreeNode[T]) tree.TreeNode[T] {
	return &BinaryTreeNode[T]{
		children: []*BinaryTreeNode[T]{left, right},
		data:     data,
	}
}

func (node BinaryTreeNode[T]) Children() []tree.TreeNode[T] {
	children := make([]tree.TreeNode[T], len(node.children))
	for i, child := range node.children {
		children[i] = child
	}
	return children
}

func (node *BinaryTreeNode[T]) SetChildren(newChildren []tree.TreeNode[T]) {
	binaryTreeNodeChildren := make([]*BinaryTreeNode[T], len(newChildren))

	for i, child := range newChildren {
		childNode := child
		binaryTreeNodeChildren[i] = &BinaryTreeNode[T]{
			data: childNode.Data(),
		}

		if len(childNode.Children()) > 0 {
			binaryTreeNodeChildren[i].SetChildren(childNode.Children())
		}
	}

	node.children = binaryTreeNodeChildren
}

func (node BinaryTreeNode[T]) Data() T {
	return node.data
}

func (node *BinaryTreeNode[T]) SetData(data T) {
	node.data = data
}

func (node BinaryTreeNode[T]) IsLeaf() bool {
	for _, child := range node.Children() {
		if child != nil {
			return false
		}
	}

	return true
}

func Left[T constraints.Ordered](node tree.TreeNode[T]) tree.TreeNode[T] {
	if node == nil ||
		node.Children() == nil ||
		len(node.Children()) == 0 {
		return nil
	}
	return node.Children()[0]
}

func Right[T constraints.Ordered](node tree.TreeNode[T]) tree.TreeNode[T] {
	if node == nil ||
		node.Children() == nil ||
		len(node.Children()) == 0 {
		return nil
	}
	return node.Children()[1]
}

func (node *BinaryTreeNode[T]) Right() *BinaryTreeNode[T] {
	return node.children[0]
}

func (node *BinaryTreeNode[T]) Left() *BinaryTreeNode[T] {
	return node.children[1]
}

func (node *BinaryTreeNode[T]) SetRight(rightNode *BinaryTreeNode[T]) {
	node.children[0] = rightNode
}

func (node *BinaryTreeNode[T]) SetLeft(leftNode *BinaryTreeNode[T]) {
	node.children[1] = leftNode
}
