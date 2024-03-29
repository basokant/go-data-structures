package binarytree

import (
	"github.com/basokant/go-data-structures/queue"
	"golang.org/x/exp/constraints"
)

type Order int

const (
	Inorder Order = iota
	Preorder
	Postorder
	Levelorder
)

func (o Order) String() string {
	return []string{"Inorder", "Preorder", "Postorder", "Levelorder"}[o]
}

func (o Order) EnumIndex() int {
	return int(o)
}

func inorderTraversal[T constraints.Ordered](node *BinaryTreeNode[T], op func(node *BinaryTreeNode[T]) bool) bool {
	left, right := node.Left, node.Right
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

func preorderTraversal[T constraints.Ordered](node *BinaryTreeNode[T], op func(node *BinaryTreeNode[T]) bool) bool {
	left, right := node.Left, node.Right
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

func postorderTraversal[T constraints.Ordered](node *BinaryTreeNode[T], op func(node *BinaryTreeNode[T]) bool) bool {
	left, right := node.Left, node.Right
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

func levelorderTraversal[T constraints.Ordered](node *BinaryTreeNode[T], op func(node *BinaryTreeNode[T]) bool) {
	if node == nil {
		return
	}

	queue := queue.New[*BinaryTreeNode[T]]()
	queue.Enqueue(node)

	for !queue.IsEmpty() {
		node, err := queue.Dequeue()
		if err != nil && op(node) {
			return
		}

		left, right := node.Left, node.Right
		if left != nil {
			queue.Enqueue(left)
		}

		if right != nil {
			queue.Enqueue(right)
		}
	}
}

func (tree BinaryTree[T]) Traverse(op func(node *BinaryTreeNode[T]) bool, order Order) {
	switch order {
	case Preorder:
		preorderTraversal(tree.root, op)
	case Postorder:
		postorderTraversal[T](tree.root, op)
	case Levelorder:
		levelorderTraversal[T](tree.root, op)
	default:
		inorderTraversal(tree.root, op)
	}
}
