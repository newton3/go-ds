package bst

type node struct {
	key   int
	left  *node
	right *node
}

//Tree ...
type Tree struct {
	root *node
	size int
}

//New ...
func New() *Tree {
	return &Tree{}
}

//Empty ...
func (t *Tree) Empty() bool {
	return t.size == 0
}

//Add ... duplicates are ignored
func (t *Tree) Add(i int) {
	node := newNode(i)
	t.root = t.root.add(node)
	t.size++
}

func newNode(i int) *node {
	return &node{key: i}
}

func (r *node) add(n *node) *node {
	if r == nil {
		return n
	}

	if n.key == r.key {
		return r
	}

	if n.key < r.key {
		r.left = r.left.add(n)
	} else {
		r.right = r.right.add(n)
	}

	return r
}
