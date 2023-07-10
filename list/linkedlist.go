package list

import (
	"errors"
)

type Node[T comparable] struct {
	Data T
	next *Node[T]
}

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
	return nil
}

func (list *LinkedList[T]) Search(data T) (*Node[T], error) {
	var current = list.head

	for current != nil {
		if current.Data == data {
			return current, nil
		}
	}
	return nil, errors.New("node with data was not found in the List")
}

func NewLinkedList[T comparable]() List[T] {
	return &LinkedList[T]{
		length: 0,
		head:   &Node[T]{},
	}
}
