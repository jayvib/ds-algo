package list

type Node struct {
	value interface{}
	next  *Node
}

func New() *List {
	return &List{}
}

type List struct {
	size int
	head *Node
}

func (l *List) PushFront(value interface{}) {
	if l.head == nil {
		l.size++
		l.head = &Node{
			value: value,
		}
		return
	}

	l.head = &Node{
		value: value,
		next:  l.head,
	}
	l.size++
}

func (l *List) PushBack(value interface{}) {
	if l.head == nil {
		l.head = &Node{
			value: value,
		}
		l.size++
		return
	}

	var lastNode *Node
	for lastNode = l.head; lastNode != nil; lastNode = lastNode.next {
	}

	lastNode.next = &Node{value: value}
	l.size++
}
