package review

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

func (n *Node) Value() int {
	return n.property
}

func (n *Node) Next() *Node {
	return n.nextNode
}

func (n *Node) String() string {
	return fmt.Sprintf("value: %d", n.property)
}

type List struct {
	head *Node
}

func (l *List) AddToHead(v int) {
	l.head = &Node{property:v, nextNode: l.head}
}

func (l *List) GetHead() *Node {
	return l.head
}

func (l *List) Iterate() {
	for n := l.head; n != nil; n = n.nextNode {
		fmt.Println(n)
	}
}

func (l *List) LastNode() *Node {
	for n := l.head; n != nil; n = n.nextNode {
		if n.nextNode == nil {
			return n
		}
	}
	return nil
}