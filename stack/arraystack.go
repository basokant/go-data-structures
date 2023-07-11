package stack

import (
	"errors"
	"fmt"
)

type ArrayStack[T comparable] struct {
	slice  []T
	length int
}

func (stack ArrayStack[T]) Empty() bool {
	return stack.slice == nil || stack.length == 0
}

func (stack ArrayStack[T]) Peek() (T, error) {
	var top T
	if stack.Empty() {
		return top, errors.New("stack empty, cannot peek the top value")
	}

	top = stack.slice[0]
	return top, nil
}

func (stack *ArrayStack[T]) Pop() (T, error) {
	var top T
	if stack.Empty() {
		return top, errors.New("stack empty, cannot pop the top value")
	}

	top = stack.slice[0]
	stack.slice = stack.slice[1:]
	stack.length--
	return top, nil
}

func (stack *ArrayStack[T]) Push(data T) error {
	if stack.slice == nil {
		stack.slice = []T{}
	}

	stack.slice = append([]T{data}, stack.slice...)
	stack.length++
	return nil
}

func (stack ArrayStack[T]) Search(data T) error {
	if stack.Empty() {
		return errors.New("stack empty, cannot search an empty stack")
	}

	for _, v := range stack.slice {
		if v == data {
			return nil
		}
	}

	return fmt.Errorf("data was not found in the stack")
}

func NewArrayStack[T comparable]() Stack[T] {
	return &ArrayStack[T]{}
}
