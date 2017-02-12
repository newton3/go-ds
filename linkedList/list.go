package linkedList

type node struct {
	key  int
	next *node
}

//List ... Singly linked list
type List struct {
	head *node
	size int
}

//New ... create a new empty list
func New() *List {
	return &List{}
}

//Empty ...
func (l *List) Empty() bool {
	return l.size == 0
}

//Add ... add an item to the head of the list
func (l *List) Add(i int) {
	if l == nil {
		return
	}

	node := newNode(i)
	if !l.Empty() {
		node.next = l.head
	}

	l.head = node
	l.size++
}

func newNode(i int) *node {
	return &node{
		key: i,
	}
}
