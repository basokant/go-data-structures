package stack

import (
	"errors"
	"fmt"
)

type ArrayStack[T comparable] struct {
	slice  []T
	length int
}

func (stack ArrayStack[T]) IsEmpty() bool {
	return stack.slice == nil || stack.length == 0
}

func (stack ArrayStack[T]) Peek() (T, error) {
	var top T
	if stack.IsEmpty() {
		return top, errors.New("stack empty, cannot peek the top value")
	}

	top = stack.slice[0]
	return top, nil
}

func (stack *ArrayStack[T]) Pop() (T, error) {
	var top T
	if stack.IsEmpty() {
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

func (stack ArrayStack[T]) IndexOf(data T) (int, error) {
	if stack.IsEmpty() {
		return -1, errors.New("stack empty, cannot search an empty stack")
	}

	for i, v := range stack.slice {
		if v == data {
			return i, nil
		}
	}

	return -1, fmt.Errorf("data was not found in the stack")
}

func (stack ArrayStack[T]) Len() int {
	return stack.length
}

func NewArrayStack[T comparable]() Stack[T] {
	return &ArrayStack[T]{}
}

func NewArrayStackFromArray[T comparable](slice []T) Stack[T] {
	sliceCopy := make([]T, len(slice))
	copy(sliceCopy, slice)

	return &ArrayStack[T]{
		slice:  sliceCopy,
		length: len(sliceCopy),
	}
}
