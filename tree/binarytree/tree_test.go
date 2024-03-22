package binarytree

import (
	"reflect"
	"testing"

	"github.com/basokant/go-data-structures/tree"
)

func newSingleNodeBinaryTree() *BinaryTree[int] {
	binaryTree := &BinaryTree[int]{
		root: &BinaryTreeNode[int]{data: 1},
	}

	return binaryTree
}

func newFullBinaryTree() *BinaryTree[int] {
	left := &BinaryTreeNode[int]{data: 2}
	right := &BinaryTreeNode[int]{data: 3}

	root := NewBinaryTreeNode[int](1, left, right)

	binaryTree := &BinaryTree[int]{root: root.(*BinaryTreeNode[int])}
	return binaryTree
}

func TestBinaryTreeRoot(t *testing.T) {
	binaryTree := NewBinaryTree[int]()

	_, err := binaryTree.Root()
	if err == nil {
		t.Error("Tree with empty root should return an error")
	}

	tests := []struct {
		tree TraversableTree[int]
		name string
		want int
	}{
		{
			name: "single node in binary tree",
			tree: newSingleNodeBinaryTree(),
			want: 1,
		},
		{
			name: "full binary tree",
			tree: newFullBinaryTree(),
			want: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root, _ := tc.tree.Root()
			got := root.Data()

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}

	t.Run("empty binary tree", func(t *testing.T) {
		var want tree.TreeNode[int]

		tree := &BinaryTree[int]{}
		got, _ := tree.Root()

		if want != got {
			t.Fatalf("expected: %v, got: %v", want, got)
		}
	})
}
