package list

import "fmt"

type Node[T comparable] interface {
	String() (string, error)
}

type SinglyLinkedNode[T comparable] struct {
	Data T
	Next *SinglyLinkedNode[T]
}

type DoublyLinkedNode[T comparable] struct {
	Data     T
	Next     *DoublyLinkedNode[T]
	Previous *DoublyLinkedNode[T]
}

func (node SinglyLinkedNode[T]) String() string {
	return fmt.Sprintf("Node: %v -> %v", node.Data, node.Next.Data)
}

func (node DoublyLinkedNode[T]) String() string {
	return fmt.Sprintf("%v -> Node: %v -> %v", node.Previous.Data, node.Data, node.Next.Data)
}

type List[T comparable] interface {
	Prepend(data T) error
	Append(data T) error
	// Delete(node *Node[T]) (T, error)
	// Search(data T) (*Node[T], error)
	// Get(index int) (*Node[T], error)
	Array() []T
	String() string
	Len() int
	IsEmpty() bool
}
