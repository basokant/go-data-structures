package list

import "fmt"

type Node[T comparable] interface {
	Data() (T, error)
	String() (string, error)
}

type SinglyLinkedNode[T comparable] struct {
	Node[T]
	Data T
	next *SinglyLinkedNode[T]
}

type DoublyLinkedNode[T comparable] struct {
	Node[T]
	Data     T
	next     *SinglyLinkedNode[T]
	previous *SinglyLinkedNode[T]
}

func (node SinglyLinkedNode[T]) String() string {
	return fmt.Sprintf("Node: %v -> %v", node.Data, node.next.Data)
}

func (node DoublyLinkedNode[T]) String() string {
	return fmt.Sprintf("%v -> Node: %v -> %v", node.previous.Data, node.Data, node.next.Data)
}

type Lister[T comparable] interface {
	Prepend(data T) error
	Append(data T) error
	Delete(node *SinglyLinkedNode[T]) (T, error)
	Search(data T) (*SinglyLinkedNode[T], error)
	Array() []T
	String() string
}
