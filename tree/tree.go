package tree

type TreeNode[T comparable] interface {
	Children() ([]*TreeNode[T], error)
}

type Tree[T comparable] interface {
	Root() (*TreeNode[T], error)
	Insert(data T) error
	Delete(data T) error
	Search(data T) (*TreeNode[T], error)
	Traverse(op func(T)) error
	Height() (int, error)
	Level(node *TreeNode[T]) (int, error)
	Size() int
}
