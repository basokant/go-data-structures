package queue

import (
	"errors"
	"fmt"

	"github.com/basokant/go-data-structures/list"
)

type ListQueue[T comparable] struct {
	list   *list.DoublyLinkedList[T]
	Length int
}

func (queue *ListQueue[T]) Dequeue() (T, error) {
	var top T
	if queue.IsEmpty() {
		return top, errors.New("queue empty, cannot pop the top value")
	}

	top, err := queue.list.Delete(queue.list.Head)
	queue.Length--

	return top, err
}

func (queue ListQueue[T]) IsEmpty() bool {
	return queue.list.IsEmpty()
}

func (queue *ListQueue[T]) Enqueue(data T) error {
	queue.list.Prepend(data)
	queue.Length++

	return nil
}

func (queue ListQueue[T]) Front() (T, error) {
	var top T
	if queue.IsEmpty() {
		return top, errors.New("queue empty, cannot pop the top value")
	}

	top = queue.list.Head.Data
	return top, nil
}

func (queue ListQueue[T]) Back() (T, error) {
	var bot T
	if queue.IsEmpty() {
		return bot, errors.New("queue empty, cannot pop the top value")
	}

	bot = queue.list.Tail.Data
	return bot, nil
}

func (queue ListQueue[T]) Search(data T) error {
	if queue.IsEmpty() {
		return errors.New("queue empty, cannot search an empty queue")
	}

	_, _, err := queue.list.Search(data)
	if err != nil {
		return fmt.Errorf("data was not found in the queue")
	}

	return nil
}

func (queue ListQueue[T]) Len() int {
	return queue.Length
}

func NewListQueue[T comparable]() Queue[T] {
	return &ListQueue[T]{}
}
