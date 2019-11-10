package linkedlist

import (
	"fmt"
	log "github.com/jayvib/golog"
)

type Node struct {
	value int
	next  *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("My Value: %v! (☞ﾟ∀ﾟ)☞", n.value)
}

func New() *List {
	return &List{}
}

type List struct {
	head, tail *Node
	length     int
}

func (l *List) Len() int {
	return l.length
}

func (l *List) IsEmpty() bool {
	return l.Len() == 0
}

// PushFront describes list insertion classic operation
func (l *List) PushFront(v int) {
	newNode := &Node{value: v}
	if l.IsEmpty() {
		// the new node will become the new head and the tail
		l.head, l.tail = newNode, newNode
	} else {
		l.head = &Node{value: v, next: l.head}
	}
	l.length++
	log.Tracef("%#v", l)
}

func (l *List) PushBack(v int) {
	l.length++
	newNode := &Node{value: v}
	if l.IsEmpty() {
		l.head, l.tail = newNode, newNode
	} else {
		log.Tracef("%#v\n", l)
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *List) Remove(v int) {
	for curr, next := l.head, l.head.next; curr != nil; curr, next = next, next.next {
	}
}

func (l *List) Get(v int) (value int) {
	currIndex := 0
	for node := l.head; node != nil; node = node.next {
		if currIndex == v {
			return node.value
		}
		currIndex++
	}
	return
}

func (l *List) Find(v int) (node *Node) {
	l.Iterate(func(n *Node) bool {
		if n.value == v {
			node = n
			return true
		}
		return false
	})
	return
}

func(l *List) IsExists(v int) bool {
	found := false
	l.Iterate(func(n *Node) bool {
		log.Tracef("current node value: %v\n", n.value)
		if n.value == v {
			log.Trace("yey! I found it! \\(^_^)/")
			found = true
			return true
		}
		return false
	})
	return found
}

func (l *List) Iterate(fn func(n *Node) bool) {
	for node := l.head; node != nil; node = node.next {
		if stop := fn(node); stop {
			return
		}
	}
}

func (l *List) Print() {
	for n := l.head; n != nil; n = n.next {
		fmt.Println(n.value)
	}
}
