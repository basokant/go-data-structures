package list

type List[T comparable] interface {
	Prepend(data T) error
	Append(data T) error
	Delete(data T) (T, error)
	IndexOf(data T) (int, error)
	Get(index int) (T, error)
	Array() []T
	String() string
	Len() int
	IsEmpty() bool
}
