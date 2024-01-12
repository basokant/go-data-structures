package tree

import (
	"github.com/basokant/go-data-structures/queue"
	"golang.org/x/exp/constraints"
)

type BinaryTreeNode[T constraints.Ordered] struct {
	data     T
	children []BinaryTreeNode[T]
}

func NewBinaryTreeNode[T constraints.Ordered](left BinaryTreeNode[T], right BinaryTreeNode[T]) TreeNode[T] {
	return BinaryTreeNode[T]{
		children: []BinaryTreeNode[T]{left, right},
	}
}

type BinaryTree[T constraints.Ordered] struct {
	root BinaryTreeNode[T]
	len  int
}

func NewBinaryTree[T constraints.Ordered]() Tree[T] {
	return BinaryTree[T]{}
}

func (node BinaryTreeNode[T]) Children() []TreeNode[T] {
	children := make([]TreeNode[T], len(node.children))
	for i, child := range node.children {
		children[i] = child
	}
	return children
}

func (node BinaryTreeNode[T]) Data() T {
	return node.data
}

func (node BinaryTreeNode[T]) IsLeaf() bool {
	for _, child := range node.Children() {
		if child != nil {
			return false
		}
	}

	return true
}

func (node BinaryTreeNode[T]) Left() TreeNode[T] {
	if node.children == nil {
		return nil
	}

	return node.children[0]
}

func left[T constraints.Ordered](node TreeNode[T]) TreeNode[T] {
	if node.Children() == nil {
		return nil
	}
	return node.Children()[0]
}

func (node BinaryTreeNode[T]) Right() TreeNode[T] {
	return node.children[1]
}

func right[T constraints.Ordered](node TreeNode[T]) TreeNode[T] {
	if node.Children() == nil {
		return nil
	}
	return node.Children()[1]
}

func inorderTraversal[T constraints.Ordered](node TreeNode[T], op func(node TreeNode[T]) bool) bool {
	left, right := left(node), right(node)
	shouldContinue := true

	if left != nil {
		shouldContinue = inorderTraversal(left, op)
		if !shouldContinue {
			return false
		}
	}

	shouldContinue = op(node)
	if !shouldContinue {
		return false
	}

	if right != nil {
		shouldContinue := inorderTraversal(right, op)
		if !shouldContinue {
			return false
		}
	}

	return true
}

func preorderTraversal[T constraints.Ordered](node TreeNode[T], op func(node TreeNode[T]) bool) bool {
	left, right := left(node), right(node)

	shouldContinue := op(node)

	if !shouldContinue {
		return false
	}

	if left != nil {
		shouldContinue = preorderTraversal(left, op)
		if !shouldContinue {
			return false
		}
	}

	if right != nil {
		shouldContinue = preorderTraversal(right, op)
		if !shouldContinue {
			return false
		}
	}

	return true
}

func postorderTraversal[T constraints.Ordered](node TreeNode[T], op func(node TreeNode[T]) bool) bool {
	left, right := left(node), right(node)
	shouldContinue := true

	if left != nil {
		shouldContinue = postorderTraversal(left, op)
		if !shouldContinue {
			return false
		}
	}

	if right != nil {
		shouldContinue = postorderTraversal(right, op)
		if !shouldContinue {
			return false
		}
	}

	shouldContinue = op(node)
	return shouldContinue
}

func levelorderTraversal[T constraints.Ordered](node TreeNode[T], op func(node TreeNode[T]) bool) {
	if node == nil {
		return
	}

	queue := queue.New[TreeNode[T]]()
	queue.Enqueue(node)

	for !queue.IsEmpty() {
		node, err := queue.Dequeue()
		if err != nil {
			op(node)
		}

		left, right := left(node), right(node)
		if left != nil {
			queue.Enqueue(left)
		}

		if right != nil {
			queue.Enqueue(right)
		}
	}
}

func (tree BinaryTree[T]) Traverse(op func(node TreeNode[T]) bool, order Order) {
	switch order {
	case Inorder:
		inorderTraversal(tree.root, op)
	case Preorder:
		preorderTraversal(tree.root, op)
	case Postorder:
		postorderTraversal[T](tree.root, op)
	case Levelorder:
		levelorderTraversal[T](tree.root, op)
	}
}

func (BinaryTree[T]) Delete(data T) error {
	panic("unimplemented")
}

func (BinaryTree[T]) Height() int {
	panic("unimplemented")
}

func (BinaryTree[T]) Insert(data T) error {
	panic("unimplemented")
}

func (BinaryTree[T]) Len() int {
	panic("unimplemented")
}

func (BinaryTree[T]) Level(node TreeNode[T]) (int, error) {
	panic("unimplemented")
}

func (BinaryTree[T]) Root() (TreeNode[T], error) {
	panic("unimplemented")
}

func (BinaryTree[T]) Search(data T) (TreeNode[T], error) {
	panic("unimplemented")
}
