package tree

type BinaryTreeNode[T comparable] struct {
	data  T
	left  *BinaryTreeNode[T]
	right *BinaryTreeNode[T]
}

func NewBinaryTreeNode[T comparable]() TreeNode[T] {
	return &BinaryTreeNode[T]{}
}

// Children implements TreeNode.
func (*BinaryTreeNode[T]) Children() ([]*TreeNode[T], error) {
	panic("unimplemented")
}

// Root implements TreeNode.
func (*BinaryTreeNode[T]) Root() (*TreeNode[T], error) {
	panic("unimplemented")
}

type BinaryTree[T comparable] struct {
	root *BinaryTreeNode[T]
	size int
}

// Delete implements Tree.
func (*BinaryTree[T]) Delete(data T) error {
	panic("unimplemented")
}

// Height implements Tree.
func (*BinaryTree[T]) Height() (int, error) {
	panic("unimplemented")
}

// Insert implements Tree.
func (*BinaryTree[T]) Insert(data T) error {
	panic("unimplemented")
}

// Level implements Tree.
func (*BinaryTree[T]) Level(node *TreeNode[T]) (int, error) {
	panic("unimplemented")
}

// Search implements Tree.
func (*BinaryTree[T]) Search(data T) (*TreeNode[T], error) {
	panic("unimplemented")
}

// Size implements Tree.
func (*BinaryTree[T]) Size() int {
	panic("unimplemented")
}

// Traverse implements Tree.
func (*BinaryTree[T]) Traverse(op func(T)) error {
	panic("unimplemented")
}

func NewBinaryTree[T comparable]() Tree[T] {
	return &BinaryTree[T]{}
}
