package stack

import (
	"errors"
	"fmt"

	"github.com/basokant/go-data-structures/list"
)

type ListStack[T comparable] struct {
	list   *list.DoublyLinkedList[T]
	length int
}

func (stack ListStack[T]) IsEmpty() bool {
	return stack.list == nil || stack.length == 0
}

func (stack ListStack[T]) Peek() (T, error) {
	var top T

	if stack.IsEmpty() {
		return top, errors.New("stack empty, cannot peek the top value")
	}

	node, err := stack.list.Get(0)
	if err != nil {
		return top, errors.Join(errors.New("top not found"), err)
	}

	top = node.Data
	return top, nil
}

func (stack *ListStack[T]) Pop() (T, error) {
	var top T

	if stack.IsEmpty() {
		return top, errors.New("stack empty, cannot peek the top value")
	}

	node, err := stack.list.Get(0)
	if err != nil {
		return top, errors.Join(errors.New("top not found"), err)
	}

	top, err = stack.list.Delete(node)
	if err != nil {
		return top, errors.Join(errors.New("top could not be deleted"), err)
	}

	stack.length--
	return top, nil
}

func (stack *ListStack[T]) Push(data T) error {
	if stack.IsEmpty() {
		stack.list = list.NewDoublyLinkedList[T]()
	}

	stack.list.Prepend(data)
	stack.length++
	return nil
}

func (stack ListStack[T]) Search(data T) (int, error) {
	index, _, err := stack.list.Search(data)

	if err != nil {
		return -1, errors.Join(fmt.Errorf("data was not found in the stack"), err)
	}

	return index, nil
}

func (stack ListStack[T]) Len() int {
	return stack.length
}

func (stack ListStack[T]) Empty() bool {
	return stack.list.IsEmpty() || stack.length == 0
}

func NewListStack[T comparable]() Stack[T] {
	return &ListStack[T]{
		list:   list.NewDoublyLinkedList[T](),
		length: 0,
	}
}

func NewListStackFromArray[T comparable](arr []T) Stack[T] {
	linkedList := list.NewDoublyLinkedListFromArray[T](arr)
	return &ListStack[T]{
		list:   linkedList,
		length: linkedList.Len(),
	}
}
