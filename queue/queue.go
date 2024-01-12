package queue

type Queue[T comparable] interface {
	Enqueue(data T) error
	Dequeue() (T, error)
	Front() (T, error)
	Search(data T) error
	IsEmpty() bool
	Len() int
}
