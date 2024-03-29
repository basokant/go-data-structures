package queue

import (
	"errors"
	"fmt"
)

type ArrayQueue[T comparable] struct {
	slice  []T
	length int
}

// Dequeue implements Queue.
func (queue *ArrayQueue[T]) Dequeue() (T, error) {
	var top T
	if queue.IsEmpty() {
		return top, errors.New("queue empty, cannot pop the top value")
	}

	top = queue.slice[0]
	queue.slice = queue.slice[1:]
	queue.length--
	return top, nil
}

// IsEmpty implements Queue.
func (queue ArrayQueue[T]) IsEmpty() bool {
	return queue.slice == nil || queue.length == 0
}

// Enqueue implements Queue.
func (queue *ArrayQueue[T]) Enqueue(data T) error {
	if queue.slice == nil {
		queue.slice = []T{}
	}

	queue.slice = append(queue.slice, data)
	queue.length++
	return nil
}

// Front implements Queue.
func (queue ArrayQueue[T]) Front() (T, error) {
	var top T
	if queue.IsEmpty() {
		return top, errors.New("queue empty, cannot pop the top value")
	}

	top = queue.slice[0]
	return top, nil
}

// Search implements Queue.
func (queue ArrayQueue[T]) Search(data T) error {
	if queue.IsEmpty() {
		return errors.New("queue empty, cannot search an empty queue")
	}

	for _, v := range queue.slice {
		if v == data {
			return nil
		}
	}

	return fmt.Errorf("data was not found in the queue")
}

func (queue ArrayQueue[T]) Len() int {
	return queue.length
}

func New[T comparable]() Queue[T] {
	return &ArrayQueue[T]{}
}

func NewFromArray[T comparable](slice []T) Queue[T] {
	sliceCopy := make([]T, len(slice))
	copy(sliceCopy, slice)

	return &ArrayQueue[T]{
		slice:  sliceCopy,
		length: len(sliceCopy),
	}
}
