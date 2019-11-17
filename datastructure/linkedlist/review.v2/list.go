package main

import (
	"fmt"
)

type Node struct {
	property int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) AddToHead(v int) {
	l.head = &Node{property: v, next: l.head}
}

func (l *LinkedList) Iterate(fn func(n *Node) (stop bool)) {
	for n := l.head; n != nil; n = n.next {
		if stop := fn(n); stop {
			break
		}
	}
}

func (l *LinkedList) Print() {
	l.Iterate(func(n *Node)bool{
		fmt.Println(n.property)
		return false
	})
}

func (l *LinkedList) LastNode() *Node {
	var last *Node
	l.Iterate(func(n *Node)bool{
		if n.next == nil {
			last = n
			return true
		}
		return false
	})
	return last
}

func (l *LinkedList) AddToEnd(v int) {
	last := l.LastNode()
	last.next = &Node{property:v}
}

func (l *LinkedList) FindNode(v int) (node *Node) {
	l.Iterate(func(n *Node)bool{
		if n.property == v {
			node = n
			return true
		}
		return false
	})
	return
}

func (l *LinkedList) AddAfter(after int, v int) {
	l.Iterate(func(n *Node)bool{
		if n.property == after {
			newNode := &Node{property: v, next: n.next}
			n.next = newNode
			return true
		}
		return false
	})
}

func main() {
	list := &LinkedList{}
	list.AddToHead(1)
	list.AddToHead(2)
	list.AddToHead(3)

	list.Iterate(func(n *Node)bool{
		fmt.Println(n.property)
		return false
	})

	list.AddToEnd(5)
	n := list.LastNode()
	fmt.Println("Last Node:", n.property)
	list.Print()

	n = list.FindNode(2)
	fmt.Print("Find node of 2:", n.property)
}
