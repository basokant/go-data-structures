package list

import "fmt"

type Node[T comparable] struct {
	Data T
	next *Node[T]
}

func (node Node[T]) String() string {
	return fmt.Sprintf("%v %v", node.Data, node.next.Data)
}

type Lister[T comparable] interface {
	Prepend(data T) error
	Append(data T) error
	Delete(node *Node[T]) error
	Search(data T) (*Node[T], error)
	Array() []T
	String() string
}
