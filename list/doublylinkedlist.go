package list

import (
	"errors"
	"fmt"
)

type DoublyLinkedList[T comparable] struct {
	Length int
	Head   *DoublyLinkedNode[T]
	Tail   *DoublyLinkedNode[T]
}

func (list *DoublyLinkedList[T]) Prepend(data T) error {
	if list == nil {
		return errors.New("cannot prepend new node on nil list")
	}

	newNode := &DoublyLinkedNode[T]{
		Data: T(data),
	}

	if list.Length == 0 {
		list.Head = newNode
		list.Tail = newNode
	} else {
		temp := list.Head
		newNode.Next = temp
		list.Head = newNode
	}

	list.Length += 1
	return nil
}

func (list *DoublyLinkedList[T]) Append(data T) error {
	if list == nil {
		return errors.New("cannot prepend new node on nil list")
	}

	newNode := &DoublyLinkedNode[T]{
		Data: T(data),
	}

	if list.Length == 0 {
		list.Head = newNode
		list.Tail = newNode
	} else {
		temp := list.Tail
		temp.Next = newNode
		list.Tail = newNode
	}
	list.Length += 1
	return nil
}

func (list *DoublyLinkedList[T]) Delete(node *DoublyLinkedNode[T]) (T, error) {
	var data T
	if list == nil {
		return data, errors.New("cannot delete from a nil list")
	}

	if list.Length == 1 {
		data = list.Head.Data
		list.Head = nil
		list.Tail = nil
		return data, nil
	}

	if list.Head == node {
		data = list.Head.Data
		list.Head = list.Head.Next
		return data, nil
	}

	if list.Tail == node {
		data = list.Tail.Data
		temp := list.Tail.Previous
		temp.Next = nil
		list.Tail = temp
		return data, nil
	}
	current := list.Head.Next

	for current != nil {
		if current == node {
			// node is found: delete it
			data = current.Data
			current.Previous.Next = current.Next
			return data, nil
		}
		current = current.Next
	}

	return data, errors.New("node with data was not found in the List")
}

func (list DoublyLinkedList[T]) Search(data T) (*DoublyLinkedNode[T], error) {
	current := list.Head

	for current != nil {
		if current.Data == data {
			return current, nil
		}
		current = current.Next
	}
	return nil, errors.New("node with data was not found in the List")
}

func (list DoublyLinkedList[T]) Get(index int) (*DoublyLinkedNode[T], error) {
	current := list.Head

	if index >= list.Length {
		return nil, errors.New("index out of bounds of the List")
	}

	for i := 1; current != nil && i < index; i++ {
		current = current.Next
	}

	return current, nil
}

func (list DoublyLinkedList[T]) Array() []T {
	arr := []T{}
	current := list.Head

	for current != nil {
		arr = append(arr, current.Data)
		current = current.Next
	}

	return arr
}

func (list DoublyLinkedList[T]) String() string {
	return fmt.Sprint(list.Array())
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		Length: 0,
		Head:   &DoublyLinkedNode[T]{},
	}
}

func NewDoublyLinkedListFromArray[T comparable](arr []T) *DoublyLinkedList[T] {
	list := &DoublyLinkedList[T]{}
	for _, data := range arr {
		list.Append(data)
	}
	return list
}
