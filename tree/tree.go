package tree

type TreeNode[T comparable] interface {
	Root() (*TreeNode[T], error)
	Children() ([]*TreeNode[T], error)
}

type Tree[T comparable] interface {
	Insert(data T) error
	Delete(data T) error
	Search(data T) (*TreeNode[T], error)
	Traverse(op func(T)) error
	Height() (int, error)
	Level(node *TreeNode[T]) (int, error)
	Size() int
}
