package single_linked_list

import (
	"errors"
	"fmt"
)

// Inserting:
// - At the front of the linked list
// - After a given node.
// - At the end of the linked list.

type Node struct {
	data interface{}
	next *Node
}

func (n *Node) Value() interface{} {
	return n.data
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *LinkedList) PushFront(data interface{}) {
	l.head = &Node{
		data: data,
		next: l.head,
	}
	if l.IsEmpty() {
		l.tail = l.head
	}
	l.size++
}

func (l *LinkedList) PushBack(data interface{}) {
	newNode := &Node{
		data: data,
	}
	if l.IsEmpty() {
		l.head = newNode
		l.tail = l.head
	} else {
		// Find the tail
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

func (l *LinkedList) IsEmpty() bool {
	return l.Len() == 0
}

func (l *LinkedList) Print() {
	l.traverse(func(i int, n *Node)bool{
		fmt.Print(n.data, " ")
		return false
	})
}

var ErrIndexBound = errors.New("index is beyond the current size")

func (l *LinkedList) InsertAt(index int, data interface{}) error {

	if l.Len() < index {
		return ErrIndexBound
	}

	if index == 0 || l.IsEmpty() {
		l.PushFront(data)
		return nil
	}

	if index == -1 {
		l.PushBack(data)
		return nil
	}

	l.insertAt(data, index)
	return nil
}

func (l *LinkedList) insertAt(data interface{}, index int) {
	newNode := &Node{
		data: data,
	}
	// Find the exiting node at index i
	// Let the value of index i be the new Node
	// set the current node be the next node of the new Node
	selectedNode := l.findNodeBefore(index)
	newNode.next = selectedNode.next
	selectedNode.next = newNode
	l.size++
}

func (l *LinkedList) findNodeBefore(index int) *Node {
	selectedNode := l.head
	for i := 1; i != index && i < l.Len(); i++ {
		selectedNode = selectedNode.next
	}
	return selectedNode
}

func (l *LinkedList) Len() int {
	return l.size
}

func (l *LinkedList) Get(index int) (value interface{}, err error) {
	if l.Len()-1 < index {
		return 0, ErrIndexBound
	}
	return l.getNode(index).data, nil
}

func (l *LinkedList) getNode(index int) *Node {
	var node *Node
	l.traverse(func(i int, n *Node) bool {
		if i == index {
			node = n
			return true
		}
		return false
	})
	return node
}

type TraverseHandler func(index int, node *Node) (isStop bool)

func (l *LinkedList) traverse(fn TraverseHandler) {
	for i, node := 0, l.head; i < l.Len(); i++ {
		if fn(i, node) {
			return
		}
		node = node.next
	}
}

func (l *LinkedList) Traverse(fn TraverseHandler) {
	l.traverse(fn)
}

func (l *LinkedList) Search(value int) (index int) {
	l.traverse(func(i int, node *Node)bool{
		if node.data == value {
			index = i
			return true
		}
		return false
	})
	return
}

func (l *LinkedList) Delete(index int) error {
	if l.Len()-1 < index {
		return ErrIndexBound
	}
	node := l.findNodeBefore(index)
	nextNextNode := node.next.next
	node.next = nextNextNode
	if index == l.Len()-1 {
		l.tail = node
	}
	l.size--

	return nil
}

