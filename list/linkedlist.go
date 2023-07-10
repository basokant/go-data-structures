package list

import (
	"errors"
)

type LinkedList[T comparable] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func (list *LinkedList[T]) Prepend(data T) error {
	if list == nil {
		return errors.New("cannot prepend new node on nil list")
	}

	newNode := &Node[T]{
		Data: T(data),
	}

	if list.length == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		temp := list.head
		newNode.next = temp
		list.head = newNode
	}

	list.length += 1
	return nil
}

func (list *LinkedList[T]) Append(data T) error {
	if list == nil {
		return errors.New("cannot prepend new node on nil list")
	}

	newNode := &Node[T]{
		Data: T(data),
	}

	if list.length == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		temp := list.tail
		temp.next = newNode
		list.tail = newNode
	}
	list.length += 1
	return nil
}

func (list *LinkedList[T]) Delete(node *Node[T]) error {
	if list == nil {
		return errors.New("cannot delete from a nil list")
	}

	if list.head == node {
		list.head = list.head.next
		return nil
	}

	previous := list.head
	current := list.head.next

	for current != nil {
		if current == node {
			// node is found: delete it
			previous.next = current.next
			return nil
		}

		// node was not found, traverse to the next node
		previous = current
		current = current.next
	}

	return errors.New("node with data was not found in the List")
}

func (list *LinkedList[T]) Search(data T) (*Node[T], error) {
	if list == nil {
		return nil, errors.New("cannot search a nil list")
	}
	current := list.head

	for current != nil {
		if current.Data == data {
			return current, nil
		}
		current = current.next
	}
	return nil, errors.New("node with data was not found in the List")
}

func NewLinkedList[T comparable]() List[T] {
	return &LinkedList[T]{
		length: 0,
		head:   &Node[T]{},
	}
}
