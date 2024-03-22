package binarytree

import (
	"fmt"

	"github.com/basokant/go-data-structures/tree"
)

func (bt *BinaryTree[T]) Insert(data T) error {
	inserted := false

	var empty *BinaryTreeNode[T]

	if bt.root == empty {
		bt.root = &BinaryTreeNode[T]{data: data, children: []*BinaryTreeNode[T]{}}
		inserted = true
		return nil
	}

	newNode := NewBinaryTreeNode[T](data, nil, nil)
	bt.Traverse(func(node tree.TreeNode[T]) bool {
		if inserted {
			return false
		}

		left, right := Left(node), Right(node)
		if left != nil && right != nil {
			return true
		}

		if left == nil {
			node.SetChildren([]tree.TreeNode[T]{newNode, right})
		} else if right == nil {
			node.SetChildren([]tree.TreeNode[T]{newNode, right})
		}

		inserted = true
		return false
	}, Preorder)

	if !inserted {
		return fmt.Errorf("unable to insert node with data %+v\ninto the tree", data)
	}

	return nil
}
