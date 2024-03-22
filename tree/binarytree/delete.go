package binarytree

import (
	"errors"

	"github.com/basokant/go-data-structures/tree"
)

func (bt *BinaryTree[T]) Delete(data T) error {
	deleted := false
	if bt.root == nil {
		// TODO: write error message
		return errors.New("")
	}

	if bt.root.data == data {
		r, l := bt.root.Right(), bt.root.Left()
		bt.root = nil

		switch true {
		case r != nil:
			bt.root = r
			return nil
		case l != nil:
			bt.root = l
			return nil
		}
	}

	var parentOfLastNode tree.TreeNode[T]
	direction := "right"
	lastNode, _ := bt.Root()
	left, right := Left(lastNode), Right(lastNode)
	for left == nil || right == nil {
		if right != nil {
			parentOfLastNode, lastNode = lastNode, right
			direction = "right"
		} else if left != nil {
			parentOfLastNode, lastNode = lastNode, left
			direction = "left"
		}
		left, right = Left(lastNode), Right(lastNode)
	}

	switch direction {
	case "right":
		left = Left(parentOfLastNode)
		parentOfLastNode.SetChildren([]tree.TreeNode[T]{left, nil})
	case "left":
		right = Right(parentOfLastNode)
		parentOfLastNode.SetChildren([]tree.TreeNode[T]{nil, right})
	}

	bt.Traverse(func(node tree.TreeNode[T]) bool {
		left, right := Left(node), Right(node)
		if deleted {
			return false
		}

		if left.Data() == data {
			left.SetData(lastNode.Data())
			deleted = true
			return false
		} else if right.Data() == data {
			right.SetData(lastNode.Data())
			deleted = true
			return false
		}

		return true
	}, Preorder)

	if !deleted {
		// TODO: write error message
		return errors.New("")
	}

	return nil
}
