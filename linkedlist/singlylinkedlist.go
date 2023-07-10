package list

import (
	"errors"
	"fmt"
)

type LinkedList[T comparable] struct {
	length int
	head   *SinglyLinkedNode[T]
	tail   *SinglyLinkedNode[T]
}

func (list *LinkedList[T]) Prepend(data T) error {
	if list == nil {
		return errors.New("cannot prepend new node on nil list")
	}

	newNode := &SinglyLinkedNode[T]{
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

	newNode := &SinglyLinkedNode[T]{
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

func (list *LinkedList[T]) Delete(node *SinglyLinkedNode[T]) (T, error) {
	var data T
	if list == nil {
		return data, errors.New("cannot delete from a nil list")
	}

	if list.length == 1 {
		data = list.head.Data
		list.head = nil
		list.tail = nil
		return data, nil
	}

	if list.head == node {
		data = list.head.Data
		list.head = list.head.next
		return data, nil
	}

	if list.tail == node {
		data = list.tail.Data
		current := list.head

		for current != nil && current.next != list.tail {
			current = current.next
		}

		current.next = nil
		list.tail = current
		return data, nil
	}

	previous := list.head
	current := list.head.next

	for current != nil {
		if current == node {
			// node is found: delete it
			data = current.Data
			previous.next = current.next
			return data, nil
		}

		// node was not found, traverse to the next node
		previous = current
		current = current.next
	}

	return data, errors.New("node with data was not found in the List")
}

func (list LinkedList[T]) Search(data T) (*SinglyLinkedNode[T], error) {
	current := list.head

	for current != nil {
		if current.Data == data {
			return current, nil
		}
		current = current.next
	}
	return nil, errors.New("node with data was not found in the List")
}

func (list LinkedList[T]) Array() []T {
	arr := []T{}
	current := list.head

	for current != nil {
		arr = append(arr, current.Data)
		current = current.next
	}

	return arr
}

func (list LinkedList[T]) String() string {
	return fmt.Sprint(list.Array())
}

func NewLinkedList[T comparable]() Lister[T] {
	return &LinkedList[T]{
		length: 0,
		head:   &SinglyLinkedNode[T]{},
	}
}

func NewLinkedListFromArray[T comparable](arr []T) *LinkedList[T] {
	list := &LinkedList[T]{}
	for _, data := range arr {
		list.Append(data)
	}
	return list
}
