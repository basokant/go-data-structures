package tree

type Node[T comparable] interface {
	Root() (*Node[T], error)
	Left() (*Node[T], error)
	Right() (*Node[T], error)
}

type Tree[T comparable] interface {
	Insert(data T) error
	Delete(data T) error
	Search(data T) (*Node[T], error)
	Traverse(op func(T)) error
	Height() (int, error)
	Level(node *Node[T]) (int, error)
	Size() int
}
