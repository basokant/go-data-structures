package binarytree

import "fmt"

func (bt *BinaryTree[T]) Insert(data T) error {
	inserted := false

	var empty *BinaryTreeNode[T]

	if bt.root == empty {
		bt.root = &BinaryTreeNode[T]{data: data}
		inserted = true
		return nil
	}

	newNode := &BinaryTreeNode[T]{data, nil, nil}
	bt.Traverse(func(node *BinaryTreeNode[T]) bool {
		if inserted {
			return false
		}

		left, right := node.Left, node.Right
		if left != nil && right != nil {
			return true
		}

		if left == nil {
			node.Left = newNode
		} else if right == nil {
			node.Right = newNode
		}

		inserted = true
		return false
	}, Preorder)

	if !inserted {
		return fmt.Errorf("unable to insert node with data %+v\ninto the tree", data)
	}

	return nil
}
