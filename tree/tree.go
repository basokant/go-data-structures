package tree

type TreeNode[T comparable] interface {
	Data() T
	// the tree could have any number of children (not just 2)
	Children() []TreeNode[T]
	IsLeaf() bool
}

type Order int

const (
	Inorder Order = iota
	Preorder
	Postorder
	Levelorder
)

func (o Order) String() string {
	return []string{"Inorder", "Preorder", "Postorder", "Levelorder"}[o]
}

func (o Order) EnumIndex() int {
	return int(o)
}

type Tree[T comparable] interface {
	Root() (TreeNode[T], error)
	Insert(data T) error
	Delete(data T) error
	Search(data T) (TreeNode[T], error)
	Traverse(op func(node TreeNode[T]) bool, order Order)
	Height() int
	Level(node TreeNode[T]) (int, error)
	Len() int
}
