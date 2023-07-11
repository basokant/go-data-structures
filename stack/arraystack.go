package stack

import (
	"errors"
	"fmt"
)

type ArrayStack[T comparable] struct {
	array  []T
	length int
}

func (stack *ArrayStack[T]) Empty() bool {
	return stack.length == 0
}

func (stack *ArrayStack[T]) Peek() (T, error) {
	var top T
	if stack.array == nil || stack.length == 0 {
		return top, errors.New("stack empty, cannot peek the top value")
	}

	top = stack.array[0]
	return top, nil
}

func (stack *ArrayStack[T]) Pop() (T, error) {
	var top T
	if stack.array == nil || stack.length == 0 {
		return top, errors.New("stack empty, cannot pop the top value")
	}

	top = stack.array[0]
	stack.array = stack.array[1:]
	stack.length--
	return top, nil
}

func (stack *ArrayStack[T]) Push(data T) error {
	if stack.array == nil {
		stack.array = []T{}
	}

	stack.array = append(stack.array, data)
	return nil
}

func (stack *ArrayStack[T]) Search(data T) error {
	if stack.array == nil || stack.length == 0 {
		return errors.New("stack empty, cannot search an empty stack")
	}

	for _, v := range stack.array {
		if v == data {
			return nil
		}
	}

	return fmt.Errorf("could not find %v in the stack %", data)
}

func NewArrayStack[T comparable]() Stack[T] {
	return &ArrayStack[T]{}
}
