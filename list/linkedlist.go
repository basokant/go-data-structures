package list

import (
	"errors"
	"fmt"
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

func (list LinkedList[T]) Search(data T) (*Node[T], error) {
	current := list.head

	for current != nil {
		if current.Data == data {
			return current, nil
		}
		current = current.next
	}
	return nil, errors.New("node with data was not found in the List")
}

func (list LinkedList[T]) ToArray() []T {
	arr := []T{}
	current := list.head

	for current != nil {
		arr = append(arr, current.Data)
		current = current.next
	}

	return arr
}

func (list LinkedList[T]) ToString() string {
	return fmt.Sprint(list.ToArray())
}

func NewLinkedList[T comparable]() List[T] {
	return &LinkedList[T]{
		length: 0,
		head:   &Node[T]{},
	}
}

func NewLinkedListFromArray[T comparable](arr []T) List[T] {
	list := &LinkedList[T]{}
	list.head = &Node[T]{}
	var previous *Node[T] = nil
	current := list.head

	for _, data := range arr {
		current.Data = data
		current.next = &Node[T]{}
		previous = current
		current = current.next
	}

	previous.next = nil

	return list
}
