package tree

type Tree[T comparable] interface {
	Insert(data T) error
	Delete(data T) error
	Search(data T)
}
