package tree

import (
	"errors"
	"fmt"
	"math"

	"github.com/basokant/go-data-structures/queue"
	"golang.org/x/exp/constraints"
)

type BinarySearchTreeNode[T constraints.Ordered] struct {
	data  T
	Left  *BinarySearchTreeNode[T]
	Right *BinarySearchTreeNode[T]
}

func NewBinarySearchTreeNode[T constraints.Ordered]() TreeNode[T] {
	return &BinarySearchTreeNode[T]{}
}

func (node BinarySearchTreeNode[T]) Data() T {
	return node.data
}

// Children implements TreeNode.
func (node BinarySearchTreeNode[T]) Children() ([]TreeNode[T], error) {
	if node.Left == nil && node.Right == nil {
		return nil, nil
	}
	left, right := TreeNode[T](node.Left), TreeNode[T](node.Right)

	children := make([]TreeNode[T], 2)
	children[0], children[1] = left, right

	return children, nil
}

func (node BinarySearchTreeNode[T]) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}

func (node BinarySearchTreeNode[T]) Height() int {
	if node.IsLeaf() {
		return 1
	}
	leftHeight, rightHeight := float64(node.Left.Height()), float64(node.Right.Height())

	return int(math.Max(leftHeight, rightHeight)) + 1
}

func (node BinarySearchTreeNode[T]) inorderTraverse(op func(node TreeNode[T])) {
	if node.Left != nil {
		node.Left.inorderTraverse(op)
	}

	op(node)

	if node.Right != nil {
		node.Right.inorderTraverse(op)
	}
}

func (node BinarySearchTreeNode[T]) preorderTraverse(op func(node TreeNode[T])) {
	op(node)
	if node.Left != nil {
		node.Left.inorderTraverse(op)
	}

	if node.Right != nil {
		node.Right.inorderTraverse(op)
	}
}

func (node BinarySearchTreeNode[T]) postorderTraverse(op func(node TreeNode[T])) {
	if node.Left != nil {
		node.Left.inorderTraverse(op)
	}

	if node.Right != nil {
		node.Right.inorderTraverse(op)
	}
	op(node)
}

type BinarySearchTree[T constraints.Ordered] struct {
	root *BinarySearchTreeNode[T]
	size int
}

// Root implements TreeNode.
func (tree BinarySearchTree[T]) Root() (TreeNode[T], error) {
	if tree.root == nil {
		return nil, errors.New("tree is empty, it has no root")
	}
	return tree.root, nil
}

// Insert implements Tree.
func (tree *BinarySearchTree[T]) Insert(data T) error {
	if tree == nil {
		return errors.New("cannot insert into nil tree")
	}

	node := &BinarySearchTreeNode[T]{
		data:  data,
		Left:  nil,
		Right: nil,
	}

	if tree.root == nil {
		tree.root = node
		return nil
	}

	terminalNode := tree.root

	for !terminalNode.IsLeaf() {
		if data <= terminalNode.data {
			terminalNode = terminalNode.Left
		} else if data > terminalNode.data {
			terminalNode = terminalNode.Right
		}
	}

	if data <= terminalNode.data {
		terminalNode.Left = node
	} else if data > terminalNode.data {
		terminalNode.Right = node
	}

	return nil
}

// Delete implements Tree.
func (*BinarySearchTree[T]) Delete(data T) error {
	panic("unimplemented")
}

// Height implements Tree.
func (tree BinarySearchTree[T]) Height() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.Height()
}

// Level implements Tree.
func (tree BinarySearchTree[T]) Level(node TreeNode[T]) (int, error) {
	curr, level := tree.root, 0

	if node == nil {
		return -1, errors.New("nil node does not exist in tree")
	}

	for curr != nil {
		if curr.Data() == node.Data() {
			return level, nil
		} else if node.Data() < curr.Data() {
			curr = curr.Left
			level++
		} else if node.Data() > curr.Data() {
			curr = curr.Right
			level++
		}
	}

	return -1, errors.New("could not find node with data")
}

// Search implements Tree.
func (tree BinarySearchTree[T]) Search(data T) (TreeNode[T], error) {
	node := tree.root

	for node != nil {
		if node.Data() == data {
			return node, nil
		} else if data < node.Data() {
			node = node.Left
		} else if data > node.Data() {
			node = node.Right
		}
	}

	return nil, errors.New("could not find node with data")
}

// Size implements Tree.
func (tree BinarySearchTree[T]) Len() int {
	return tree.size
}

func (tree BinarySearchTree[T]) levelorderTraverse(op func(node TreeNode[T])) {
	q := queue.NewArrayQueue[*BinarySearchTreeNode[T]]()
	q.Enqueue(tree.root)

	level := 0
	for !q.IsEmpty() {
		fmt.Println(level)
		for i := 0; i < q.Size(); i++ {
			curr, _ := q.Dequeue()
			if curr != nil {
				op(curr)
				q.Enqueue(curr.Left)
				q.Enqueue(curr.Right)
			}
		}
	}
}

// Traverse implements Tree.
func (tree BinarySearchTree[T]) Traverse(op func(node TreeNode[T]), order Order) {
	switch order {
	case Inorder:
		tree.root.inorderTraverse(op)
	case Preorder:
		tree.root.preorderTraverse(op)
	case Postorder:
		tree.root.postorderTraverse(op)
	case Levelorder:
		tree.levelorderTraverse(op)
	}
}

func NewBinarySearchTree[T constraints.Ordered]() Tree[T] {
	return &BinarySearchTree[T]{}
}
