package stack

type ListStack[T comparable] struct {
	list   []T
	length int
}

// Peek implements Stack
func (*ListStack[T]) Peek() (T, error) {
	panic("unimplemented")
}

// Pop implements Stack
func (*ListStack[T]) Pop() (T, error) {
	panic("unimplemented")
}

// Push implements Stack
func (*ListStack[T]) Push(data T) error {
	panic("unimplemented")
}

// Search implements Stack
func (*ListStack[T]) Search(data T) error {
	panic("unimplemented")
}

func (stack *ListStack[T]) Empty() bool {
	return stack.length == 0
}

func NewListStack[T comparable]() Stack[T] {
	return &ListStack[T]{}
}
