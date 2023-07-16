package tree

type TreeNode[T comparable] interface {
	Data() T
	Children() ([]TreeNode[T], error)
	IsLeaf() bool
}

type Order int

const (
	Inorder Order = iota
	Preorder
	Postorder
	Levelorder
)

// String - Creating common behavior - give the type a String function
func (o Order) String() string {
	return [...]string{"Inorder", "Preorder", "Postorder", "Levelorder"}[o]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (o Order) EnumIndex() int {
	return int(o)
}

type Tree[T comparable] interface {
	Root() (TreeNode[T], error)
	Insert(data T) error
	Delete(data T) error
	Search(data T) (TreeNode[T], error)
	Traverse(op func(node TreeNode[T]), order Order)
	Height() int
	Level(node TreeNode[T]) (int, error)
	Size() int
}
