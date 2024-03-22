package list

import "fmt"

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
