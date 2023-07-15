package tree

type TreeNode[T comparable] interface {
	Data() T
	Children() ([]TreeNode[T], error)
	IsLeaf() bool
}

type Tree[T comparable] interface {
	Root() (TreeNode[T], error)
	Insert(data T) error
	Delete(data T) error
	Search(data T) (TreeNode[T], error)
	Traverse(op func(node TreeNode[T]))
	Height() int
	Level(node TreeNode[T]) (int, error)
	Size() int
}
