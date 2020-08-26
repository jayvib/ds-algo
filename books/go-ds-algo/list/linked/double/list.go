package double

import "fmt"

// Node represents a node of a tree
type Node struct {
	// value is the value to be store in
	// this node
	value interface{}
	// Next is the next node of this node.
	next *Node
	// Prev is the previous node of this node.
	prev *Node
}

// Value returns the value of this node.
func (n *Node) Value() interface{} {
	return n.value
}

// Next returns the next node of this node.
func (n *Node) Next() *Node {
	return n.next
}

// Previous returns the previous node of this node.
func (n *Node) Previous() *Node {
	return n.prev
}

// Cases to consider:
// 1. Null Values
// 2. Only Element
// 3. First Element
// 4. General Case
// 5. Last Element

func New() *List {
	return new(List)
}

// List represents the doubly linked list
type List struct {
	// head represents the head node of this list
	head *Node
	// tail represents the tail node of this list
	tail *Node
	size int
}

func (l *List) Size() int {
	return l.size
}

func (l *List) IsEmpty() bool {
	return l.size == 0
}

// PushFront pushes a new node with value of v to the front
// of the list. v should not nil otherwise this will panic.
func (l *List) PushFront(v interface{}) {
	if v == nil {
		panic("empty value")
	}
	// First Case: When the list is yet empty
	if l.IsEmpty() {
		newNode := &Node{
			value: v,
		}
		l.head = newNode
		l.tail = newNode
		l.size++
		return
	}

	// Second Case: When the existing node in the
	// list are the head and the tail which are the same node
	// So make the new node as the new head node. The existing
	// node will become the next node of the current node.
	// Make the new node as the previous node of the current head node
	// General Case: When the the existing list has a size >2
	newNode := &Node{
		value: v,
		next:  l.head,
	}
	l.head.prev = newNode
	l.head = newNode
	l.size++
}

func (l *List) Print() {
	l.iterate(func(n *Node) bool {
		fmt.Println(n.value)
		return false
	})
}

func (l *List) PushBack(v interface{}) {
	if l.IsEmpty() {
		l.size++
		newNode := &Node{value: v}
		l.head = newNode
		l.tail = newNode
		return
	}

	// The list has already an existing item.
	// Find the tail, and the new node will be the
	// new tail of the current list. Make the current tail
	// be the previous of the new node.
	tail := l.tail
	newNode := &Node{
		value: v,
		prev:  tail,
	}
	tail.next = newNode
	l.tail = newNode
	l.size++
}

func (l *List) RemoveFront() (value interface{}, success bool) {
	if l.IsEmpty() {
		return nil, false
	}

	node := l.head
	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size--
		return node.value, true
	}

	nextNode := l.head.next
	nextNode.prev = nil
	l.head = nextNode
	l.size--

	return node.value, true
}

func (l *List) RemoveNode(value interface{}) (success bool) {
	// Check if empty
	if l.IsEmpty() {
		return false
	}

	// Check the head
	if l.size == 1 {
		// Check directly, meaning the head and the tail is the same node
		if l.head.value == value {
			l.head = nil
			l.tail = nil
			l.size--
			return true
		}
		return false
	}

	// Find the node
	var selected *Node
	l.iterate(func(n *Node) bool {
		if n.value == value {
			selected = n
			return true
		}
		return false
	})

	// The node is not exists
	if selected == nil {
		return false
	}

	// Found at the tail
	// Check the tail
	if selected.next == nil {
		prev := selected.prev
		prev.next = nil
		l.tail = prev
		l.size--
		return true
	}

	// Check the body
	next := selected.next
	prev := selected.prev

	prev.next = next
	next.prev = prev

	l.size--

	return true
}

func (l *List) Reset() {
	l.head = nil
	l.tail = nil
	l.size = 0
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

func (l *List) iterate(fn func(n *Node) (isStop bool)) {
	for n := l.head; n != nil; n = n.next {
		isStop := fn(n)
		if isStop {
			return
		}
	}
}

func (l *List) ReverseList() {
	var curr, tmp *Node
	curr = l.head
	for curr != nil {
		// Swap the previous and the next node
		// of the current node.
		tmp = curr.next
		curr.next = curr.prev
		curr.prev = tmp

		// This is the tail which is already in-front of the list
		if curr.prev == nil {
			l.tail = l.head
			l.head = curr
			return
		}
		curr = tmp
	}
}


// List should be sorted
func (l *List) RemoveDuplicate() {
	var current *Node
	current = l.head

	for current != nil {
		// To Delete
		if current.next != nil && current.next.value == current.value {
			nextNext := current.next.next
			if nextNext != nil {
				nextNext.prev = current
			}
			current.next = nextNext
			continue
		}
		current = current.next
	}
}

// o----o----o----o----o
// |    ^    |
// <--swap--->
func (l *List) reverseList() {
	var curr, tmp *Node
	curr = l.head
	for curr != nil {
		// Swap the current next's and previous node
		tmp = curr.next
		curr.next = curr.prev
		curr.prev = tmp

		// Meaning reach the tail node, so swap the current
		// head and tail
		if tmp == nil {
			l.tail = l.head
			l.head = curr
			return
		}

		curr = tmp
	}
}

func (l *List) InsertNode(pos int, value interface{}) (success bool) {

	if l.size < pos {
		return false
	}

	// Case when the position is at the head
	if pos == 0 {
		newNode := &Node{
			value: value,
			next: l.head,
			prev: l.tail,
		}

		l.head.prev = newNode
		l.tail.next = newNode

		l.head = newNode
		return true
	}

	currPos := 0
	var prev *Node
	l.iterate(func(n *Node)bool{
		if pos == currPos+1 {
			prev = n
			return true
		}
		currPos++
		return false
	})

	newNode := &Node{value:value}
	nodeToMove := prev.next

	newNode.next = nodeToMove
	newNode.prev = prev
	nodeToMove.prev = newNode
	prev.next = newNode

	return true
}
