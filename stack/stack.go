package stack

type Stack[T comparable] interface {
	Push(data T) error
	Pop() (T, error)
	Peek() (T, error)
	Search(data T) (int, error)
	Empty() bool
	Len() int
}
