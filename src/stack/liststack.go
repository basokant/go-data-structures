package stack

import (
	"errors"
	"fmt"

	"github.com/basokant/go-data-structures/src/list"
)

type ListStack[T comparable] struct {
	list   *list.DoublyLinkedList[T]
	length int
}

func (stack *ListStack[T]) Empty() bool {
	return stack.length == 0
}

func (stack *ListStack[T]) Peek() (T, error) {
	var top T

	if stack.length == 0 {
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

	if stack.length == 0 {
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

	return top, nil
}

func (stack *ListStack[T]) Push(data T) error {
	if stack.list == nil || stack.list.Length == 0 {
		stack.list = list.NewDoublyLinkedList[T]()
	}

	stack.list.Append(data)
	stack.length++
	return nil
}

func (stack *ListStack[T]) Search(data T) error {
	_, err := stack.list.Search(data)

	if err != nil {
		return errors.Join(fmt.Errorf("data was not found in the stack"), err)
	}

	return nil
}

func NewListStack[T comparable]() Stack[T] {
	return &ListStack[T]{}
}
