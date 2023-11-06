package list

import (
	"errors"
	"fmt"
)

type SinglyLinkedList[T comparable] struct {
	Head   *SinglyLinkedNode[T]
	Tail   *SinglyLinkedNode[T]
	Length int
}

func (list *SinglyLinkedList[T]) Prepend(data T) error {
	if list == nil {
		return errors.New("cannot prepend new node on nil list")
	}

	newNode := &SinglyLinkedNode[T]{
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

func (list *SinglyLinkedList[T]) Append(data T) error {
	if list == nil {
		return errors.New("cannot prepend new node on nil list")
	}

	newNode := &SinglyLinkedNode[T]{
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

func (list *SinglyLinkedList[T]) Delete(node *SinglyLinkedNode[T]) (T, error) {
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
		current := list.Head

		for current != nil && current.Next != list.Tail {
			current = current.Next
		}

		current.Next = nil
		list.Tail = current
		return data, nil
	}

	previous := list.Head
	current := list.Head.Next

	for current != nil {
		if current == node {
			// node is found: delete it
			data = current.Data
			previous.Next = current.Next
			return data, nil
		}

		// node was not found, traverse to the next node
		previous = current
		current = current.Next
	}

	return data, errors.New("node with data was not found in the List")
}

func (list SinglyLinkedList[T]) Search(data T) (int, *SinglyLinkedNode[T], error) {
	current := list.Head
	index := 0

	for current != nil {
		if current.Data == data {
			return index, current, nil
		}
		index++
		current = current.Next
	}
	return -1, nil, errors.New("node with data was not found in the List")
}

func (list SinglyLinkedList[T]) Get(index int) (*SinglyLinkedNode[T], error) {
	current := list.Head

	if index >= list.Length {
		return nil, errors.New("index out of bounds of the List")
	}

	for i := 1; current != nil && i < index; i++ {
		current = current.Next
	}

	return current, nil
}

func (list SinglyLinkedList[T]) Array() []T {
	arr := []T{}
	current := list.Head

	for current != nil {
		arr = append(arr, current.Data)
		current = current.Next
	}

	return arr
}

func (list SinglyLinkedList[T]) String() string {
	return fmt.Sprint(list.Array())
}

func (list SinglyLinkedList[T]) Len() int {
	return list.Length
}

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		Length: 0,
		Head:   &SinglyLinkedNode[T]{},
	}
}

func NewSinglyLinkedListFromArray[T comparable](arr []T) *SinglyLinkedList[T] {
	list := &SinglyLinkedList[T]{}
	for _, data := range arr {
		list.Append(data)
	}
	return list
}
