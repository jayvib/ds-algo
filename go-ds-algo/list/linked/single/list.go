package list

import (
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

func New() *List {
	return &List{}
}

type List struct {
	head *Node
	tail *Node
	size int
}

func (l *List) PushFront(value interface{}) {
	l.size++
	l.head = &Node{
		value: value,
		next:  l.head,
	}
}

func (l *List) PushBack(value interface{}) {
	newNode := &Node{value: value}
	if l.isEmpty() {
		l.head = newNode
		l.tail = newNode
		l.size++
		return
	}

	// Find the last node
	lastNode := l.tail
	lastNode.next = newNode
	l.tail = newNode
	l.size++
}

func (l *List) FindNode(value interface{}) (*Node, bool) {
	var node *Node
	l.Iterate(func(n *Node) bool {
		if n.value == value {
			node = n
			return true
		}
		return false
	})
	if node != nil {
		return node, true
	}
	return nil, false
}

func (l *List) RemoveFront() (value interface{}, ok bool) {
	if l.isEmpty() {
		return nil, false
	}
	currentHead := l.head
	newHead := currentHead.next
	l.head = newHead
	l.size--
	return currentHead.value, true
}

func (l *List) Iterate(fn func(n *Node) bool) {
	for node := l.head; node != nil; node = node.next {
		ok := fn(node)
		if ok {
			return
		}
	}
}

func (l *List) Tail() *Node {
	return l.tail
}

// DeleteNode deletes the existing node that has a value of v.
// When the deletion is successful it will return the value associated
// with the deleted node and sets value to true for ok.
func (l *List) DeleteNode(v interface{}) (value interface{}, ok bool) {
	// Find the node to be deleted but need to track the previous
	// node of the deleted node so that it will be connect to the
	// next node of the deleted node.

	// Only the head exists.
	if l.size == 1 {
		// Check if the value matches with the value
		// in the existing node
		if l.head.value == v {
			// Remove the head
			l.head = nil
			l.tail = nil
			l.size--
			return v, true
		}
		return nil, false
	}

	var node *Node
	// Find the node to be deleted
	l.Iterate(func(n *Node) bool {
		if n.next != nil {
			if n.next.value == v {
				node = n
				return true
			}
		}
		return false
	})

	// The node is not exists
	if node == nil {
		return nil, false
	}

	nodeToDelete := node.next
	newNext := nodeToDelete.next
	node.next = newNext

	// Check if the deleted node is the tail
	if nodeToDelete == l.tail {
		l.tail = 	node
	}
	l.size--
	return nodeToDelete.value, true
}

// DeleteNodes deletes all the nodes that has the value
// of v and return true to flag as success.
func (l *List) DeleteNodes(v interface{}) (success bool) {
	for {
		_, ok := l.DeleteNode(v)
		if !ok {
			break
		}
		if !success {
			success = true
		}
	}
	return
}

func (l *List) Reset() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *List) Print() {
	l.Iterate(func(n *Node) bool {
		fmt.Println(n.value)
		return false
	})
}

// Reverse reverses the current order of the linked list
// Illustration:
//
//       o------o-----o-----o
//     / |      |
// prev curr   next
//    |  |      |
//    <---swap-->
//       |
//    new prev(new head)
func (l *List) Reverse() {
	var prev, curr, next *Node

	// Start with the head
	curr = l.head
	for curr != nil {
		// Find the next node of the current node
		next = curr.next
		// Make the previous of the current node as the next node
		// of the current node.
		curr.next = prev
		// Set the current node as the new previous node
		prev = curr
		// Move to the next node
		curr = next
	}
	l.head = prev
}

func (l *List) RemoveDuplicateNodes() {
	for curr := l.head; curr != nil; {
		if curr.next != nil && curr.value == curr.next.value {
			curr.next = curr.next.next
		} else {
			curr = curr.next
		}
	}
}

func (l *List) CopyListReversed() *List {
	var tmp, tmp2 *Node
	for curr := l.head; curr != nil; curr = curr.next {
		tmp2 = &Node{curr.value, tmp}
		tmp = tmp2
	}
	return &List{head:tmp}
}

func (l *List) Copy() *List {
	var head, tail *Node
	for curr := l.head; curr != nil; curr = curr.next {
		node := &Node{value: curr.value}
		if head == nil {
			head = node
			tail = node
			head.next = tail
		} else {
			if tail != nil {
				tail.next = node
				tail = node
			}
		}
	}
	return &List{head: head, size: l.size}
}

func (l *List) Compare(to *List) (equal bool) {
	return compare(l.head, to.head)
}

func compare(from, to *Node) bool {
	// Base Case is when the two nodes are empty
	// When it reach to this point, meaning the
	// two list have identical values.
	if from == nil && to == nil {
		return true
	}

	// When the value of the two nodes
	// are not equal then no need to
	// further checking just return false
	// immediately.
	if (from == nil || to == nil) || (from.value != to.value) {
		return  false
	}

	return compare(from.next, to.next)
}

func (l *List) removeDuplicate() {
	curr := l.head
	for curr != nil {
		if curr.next != nil && curr.value == curr.next.value {
			curr.next = curr.next.next
		} else {
			curr = curr.next
		}
	}
}

func (l *List) recursiveReverse() {
	l.head = recursiveReverse2(l.head, nil)
}

func (l *List) LoopDetect() (found bool) {
	// Approach: Use hash map
	trackingNodeAddress := make(map[*Node]struct{})

	l.Iterate(func(n *Node)bool{
		_, ok := trackingNodeAddress[n]
		if !ok {
			trackingNodeAddress[n] = struct{}{}
			return false
		}
		found = true
		return true
	})

	return
}

// Floyd's Cycle-Finding Algorithm
func (l *List) LoopDetectFloydCycle() (found bool){
	_, found = l.findNodeLoopPoint()
	return
}

func (l *List) findNodeLoopPoint() (*Node, bool) {
	// slow will be the current node
	slow := l.head
	// fast will advance to the next of next node
	fast := l.head

	for slow.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return slow, true
		}
	}
	return nil, false
}


func (l *List) LoopDetectReverse() (found bool) {
	// Store the current head
	curr := l.head

	// Reverse the link and compare if the
	// current head is the same with the new head
	// with the reversed link. If the same,
	// meaning there's a loop. Then revert back
	// the loop.
	l.Reverse()
	if curr == l.head {
		// Reverse back the link
		l.Reverse()
		return true
	}

	// Reverse back. Meaning no found loop
	l.Reverse()
	// Reverse back the link
	return false
}

func recursiveReverse(current, prev *Node) *Node {

	// Base case
	if current == nil {
		return nil
	}

	// This will be the start of reversing.
	if current.next == nil {
		// Swap with the existing prev node
		// to the prev node of the current node
		current.next = prev

		// This will become the new head
		return current
	}

	result := recursiveReverse(current.next, current)

	current.next = prev

	return result
}

func recursiveReverse2(current, prev *Node) *Node {
	// Base case
	// When the next node is empty
	if current.next == nil {
		current.next = prev
		return current
	}

	res := recursiveReverse2(current.next, current)

	current.next = prev
	return res
}

func (l *List) reverse() {
	var prev, curr, next *Node
	curr = l.head
	for curr != nil {
		// Store the next node of the current node.
		next = curr.next
		// Set the next node's new value from the previous node.
		curr.next = prev
		// A new previous ist the current node for the next node.
		prev = curr
		// Move the cursor to the next node of the current node.
		curr = next
	}
	l.head = prev
}

func (l *List) isEmpty() bool {
	return l.size == 0
}

func (l *List) RemoveLoop() (ok bool) {
	// Find the loop detect point
	looPoint, found := l.findNodeLoopPoint()
	if !found {
		return false
	}

	// Check if the loop point is the head.
	// Meaning it is a circular list.
	// Find the tail of the list.	Which is
	// before the loop point
	fp := l.head
	if looPoint == l.head {
		// Circular List
		// Find the tail.
		if fp.next != l.head {
			// This is part of the body
			// Slide to the next item
			fp = fp.next
		}

		// Tail of the node
		// so fp next node should be nil
		fp.next = nil
		return true
	}

	// This case when the loop point is not
	// circular.
	// Find the
	sp := looPoint
	if fp.next != sp.next {
		fp = fp.next
		sp = fp.next
	}
	sp.next = nil

	 return true
}
