package binarytree

import (
	"reflect"
	"testing"
)

func newSingleNodeBinaryTree() *BinaryTree[int] {
	binaryTree := &BinaryTree[int]{}

	binaryTree.Insert(1)
	return binaryTree
}

func newFullTree() *BinaryTree[int] {
	binaryTree := &BinaryTree[int]{}

	return binaryTree
}

func TestBinaryTreeRoot(t *testing.T) {
	binaryTree := NewBinaryTree[int]()

	_, err := binaryTree.Root()
	if err == nil {
		t.Error("Tree with empty root should return an error")
	}

	tests := []struct {
		name string
		tree TraversableTree[int]
		want int
	}{
		{
			name: "single node in binary tree",
			tree: newSingleNodeBinaryTree(),
			want: 1,
		},
		{
			name: "full binary tree",
			tree: nil,
			want: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := tc.tree.Root()

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
