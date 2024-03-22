package tree

type TreeNode[T comparable] interface {
	Data() T
	SetData(data T)
	Children() []TreeNode[T]
	SetChildren(newChildren []TreeNode[T])
	IsLeaf() bool
}

type Tree[T comparable] interface {
	Root() (TreeNode[T], error)
	Insert(data T) error
	Delete(data T) error
	Search(data T) (TreeNode[T], error)
	Height() int
	Level(node TreeNode[T]) (int, error)
	Len() int
}
