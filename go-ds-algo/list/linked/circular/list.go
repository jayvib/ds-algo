package circular

import (
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

func New() *List {
	return new(List)
}

type List struct {
	head *Node
	tail *Node
	size int
}

func (l *List) Size() int {
	return l.size
}

func (l *List) isEmpty() bool {
	return l.size == 0
}

func (l *List) PushFront(value interface{}) {
	newNode := &Node{value: value}
	if l.isEmpty() {
		l.size++
		// For the first element. Connect the head and the tail
		newNode.next = newNode
		l.head = newNode
		l.tail = newNode
		return
	}

	l.size++
	newNode.next = l.head
	l.tail.next = newNode
	l.head = newNode
}

func (l *List) PushBack(value interface{}) {
	if l.isEmpty() {
		l.size++
		newNode := &Node{value: value}
		newNode.next = newNode
		l.head = newNode
		l.tail = newNode
		return
	}
	l.size++
	newNode := &Node{value: value}
	tail := l.tail
	newNode.next = tail.next
	tail.next = newNode
	l.tail = newNode
}

func (l *List) RemoveHead() (success bool) {

	if l.size == 0 {
		return false
	}

	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		next := l.head.next
		l.head = next
		l.tail.next = next
	}

	l.size--

	return true
}

func (l *List) IsExists(value interface{}) (found bool) {
	l.iterate(func(n *Node) bool {
		if n.value == value {
			found = true
			return true
		}
		return false
	})
	return
}

func (l *List) iterate(fn func(n *Node) (stop bool)) {
	isLast := false
	for n := l.head; !isLast; n = n.next {
		if n == l.tail {
			isLast = true
		}
		isStop := fn(n)
		if isStop {
			break
		}
	}
}

func (l *List) Print() {
	l.iterate(func(n *Node) bool {
		fmt.Println(n.value)
		return false
	})
}

func (l *List) Reset() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *List) RemoveNode(v interface{}) (success bool) {

	// Empty List Case
	if l.isEmpty() {
		return false
	}

	// Check when the node to look at is at the head
	if l.size == 1 {
		if l.head.value == v {
			l.head = nil
			l.tail = nil
			l.size--
			return true
		}
		return false
	}

	// General Case when the size of a list is >1
	var prevNode *Node
	l.iterate(func(n *Node)bool{
		if n.next.value == v {
			prevNode = n
			return true
		}
		return false
	})

	// The node to be remove is not exists
	if prevNode == nil {
		return false
	}

	// The next node will be the node to be remove.
	//
	// Get the next of next node
	nextNext := prevNode.next.next
	prevNode.next = nextNext
	if nextNext == l.head {
		// this is the tail
		l.tail = prevNode
	}
	l.size--

	return true
}

func (l *List) Reverse() {
}

// Given that the list is sorted
func (l *List) RemoveDuplicate() (success bool) {
	if l.size == 0 || l.size == 1 {
		return false
	}

	for curr := l.head; curr != nil && curr != l.tail; {
		if curr.value == curr.next.value {
			nextNext := curr.next.next
			curr.next = nextNext
			l.size--
			continue
		}

		curr = curr.next
	}


	return true
}
