package binarytree

import "errors"

func (bt *BinaryTree[T]) Delete(data T) error {
	deleted := false
	if bt.root == nil {
		// TODO: write error message
		return errors.New("")
	}

	if bt.root.data == data {
		r, l := bt.root.Right, bt.root.Left
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

	var parentOfLastNode *BinaryTreeNode[T]
	direction := "right"
	lastNode, _ := bt.Root()
	left, right := lastNode.Left, lastNode.Right
	for left == nil || right == nil {
		if right != nil {
			parentOfLastNode, lastNode = lastNode, right
			direction = "right"
		} else if left != nil {
			parentOfLastNode, lastNode = lastNode, left
			direction = "left"
		}
		left, right = lastNode.Left, lastNode.Right
	}

	switch direction {
	case "right":
		parentOfLastNode.Right = nil
	case "left":
		parentOfLastNode.Left = nil
	}

	bt.Traverse(func(node *BinaryTreeNode[T]) bool {
		left, right := node.Left, node.Right
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
