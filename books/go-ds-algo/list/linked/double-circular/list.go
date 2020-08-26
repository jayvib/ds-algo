package double_circular

import "fmt"

type Node struct {
	value      interface{}
	next, prev *Node
}

func New() *List {
	return &List{}
}

type List struct {
	head *Node
	tail *Node
	size int
}

func (l *List) Size() int {
	return l.size
}

func (l *List) IsPresent(value interface{}) (isFound bool) {
	l.iterate(func(n *Node)bool{
		if n.value == value {
			isFound = true
			return true
		}
		return false
	})
	return
}

func (l *List) iterate(fn func(n *Node)(stop bool)) {
	for n := l.head; n != nil; n = n.next {
		isStop := fn(n)
		if isStop {
			return
		}

		if n == l.tail {
			return
		}
	}
}

func (l *List) PushBack(v interface{}) {
	newNode := &Node{
		value: v,
	}
	if l.isEmpty() {
		l.size++
		newNode.next = newNode
		newNode.prev = newNode
		l.head = newNode
		l.tail = newNode
		return
	}
	l.size++
	newNode.prev = l.tail
	newNode.next = l.head
	l.tail.next = newNode
	l.head.prev = newNode
	l.tail = newNode
}

func (l *List) Print() {
	l.iterate(func(n *Node)bool{
		fmt.Println(n.value)
		return false
	})
}

func (l *List) Reset() {
	l.head, l.tail = nil, nil
	l.size = 0
}

func (l *List) isEmpty() bool {
	return l.size == 0
}

func (l *List) PushFront(v interface{}) {

	newNode := &Node{
		value: v,
	}
	// First Case: When the list is empty
	if l.isEmpty() {
		l.size++

		// For the first element the previous and the next
		// of this node will be the node itself.
		newNode.next = newNode
		newNode.prev = newNode
		l.head = newNode
		l.tail = newNode
		return
	}

	// General Case
	l.size++

	l.head.prev = newNode // Make the new node as the previous of the current head
	newNode.next = l.head // Make the current head be the next of the new node.
	newNode.prev = l.tail // Make the current tail be the previous of the new node.
	l.head = newNode      // Make the new node as the new head.
	l.tail.next = newNode // To connect the tail to the new head node
}

func (l *List) RemoveFront() (success bool) {
	if l.isEmpty() {
		return false
	}

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size--
		return true
	}

	// General Case
	newHead := l.head.next
	newHead.prev = l.tail
	l.tail.next = newHead
	l.head = newHead
	l.size--


	return true
}
