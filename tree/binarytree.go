package tree

import (
	"errors"
	"math"

	"golang.org/x/exp/constraints"
)

type BinarySearchTreeNode[T constraints.Ordered] struct {
	Data  T
	Left  *BinarySearchTreeNode[T]
	Right *BinarySearchTreeNode[T]
}

func NewBSTNode[T constraints.Ordered]() TreeNode[T] {
	return &BinarySearchTreeNode[T]{}
}

// Children implements TreeNode.
func (node BinarySearchTreeNode[T]) Children() ([]*TreeNode[T], error) {
	if node.Left == nil && node.Right == nil {
		return nil, nil
	}
	left, right := TreeNode[T](node.Left), TreeNode[T](node.Right)

	children := make([]*TreeNode[T], 2)
	children[0], children[1] = &left, &right

	return children, nil
}

func (node BinarySearchTreeNode[T]) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}

func (node BinarySearchTreeNode[T]) height() int {
	if node.IsLeaf() {
		return 1
	}
	leftHeight, rightHeight := float64(node.Left.height()), float64(node.Right.height())

	return int(math.Max(leftHeight, rightHeight)) + 1
}

type BinarySearchTree[T constraints.Ordered] struct {
	root *BinarySearchTreeNode[T]
	size int
}

// Root implements TreeNode.
func (tree BinarySearchTree[T]) Root() (*TreeNode[T], error) {
	if tree.root == nil {
		return nil, errors.New("tree is empty, it has no root")
	}
	root := TreeNode[T](tree.root)
	return &root, nil
}

// Insert implements Tree.
func (tree *BinarySearchTree[T]) Insert(data T) error {
	if tree == nil {
		return errors.New("cannot insert into nil tree")
	}

	node := &BinarySearchTreeNode[T]{
		Data:  data,
		Left:  nil,
		Right: nil,
	}

	if tree.root == nil {
		tree.root = node
		return nil
	}

	terminalNode := tree.root

	for !terminalNode.IsLeaf() {
		if data <= terminalNode.Data {
			terminalNode = terminalNode.Left
		} else if data > terminalNode.Data {
			terminalNode = terminalNode.Right
		}
	}

	if data <= terminalNode.Data {
		terminalNode.Left = node
	} else if data > terminalNode.Data {
		terminalNode.Right = node
	}

	return nil
}

// Delete implements Tree.
func (*BinarySearchTree[T]) Delete(data T) error {
	panic("unimplemented")
}

// Height implements Tree.
func (*BinarySearchTree[T]) Height() (int, error) {
	panic("unimplemented")
}

// Level implements Tree.
func (*BinarySearchTree[T]) Level(node *TreeNode[T]) (int, error) {
	panic("unimplemented")
}

// Search implements Tree.
func (*BinarySearchTree[T]) Search(data T) (*TreeNode[T], error) {
	panic("unimplemented")
}

// Size implements Tree.
func (tree *BinarySearchTree[T]) Size() int {
	return tree.size
}

// Traverse implements Tree.
func (*BinarySearchTree[T]) Traverse(op func(T)) error {
	panic("unimplemented")
}

func NewBinaryTree[T constraints.Ordered]() Tree[T] {
	return &BinarySearchTree[T]{}
}
