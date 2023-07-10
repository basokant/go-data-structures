package list

type Node[T comparable] struct {
	Data T
	next *Node[T]
}
type List[T comparable] interface {
	Prepend(data T) error
	Append(data T) error
	Delete(node *Node[T]) error
	Search(data T) (*Node[T], error)
	ToArray() []T
	ToString() string
}
