package stack

type Stack[T comparable] interface {
	Push(data T) error
	Pop() (T, error)
	Peek() (T, error)
	IndexOf(data T) (int, error)
	IsEmpty() bool
	Len() int
}
