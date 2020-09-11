package binary_tree

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) Value() int {
	return n.value
}

type Tree struct {
	size int
	root *Node
}

func (t *Tree) Print() {
	t.InOrderTraverse(func(n *Node){
		fmt.Println(n.value)
	})
}

func (t *Tree) InOrderTraverse(fn func(n *Node)) {
	inOrderTraverse(t.root, fn)
}

func inOrderTraverse(n *Node, fn func(n *Node)) {
	if n == nil {
		return
	}
	inOrderTraverse(n.left, fn)
	fn(n)
	inOrderTraverse(n.right, fn)
}


func (t *Tree) Root() *Node {
	return t.root
}

func (t *Tree) Insert(data int) {

	if t.isEmpty() {
		t.root = &Node{
			value: data,
		}
	} else {
		insert(t.root, data)
	}

	t.size++
}

func (t *Tree) Search(x int) *Node {
	return search(t.root, x)
}

func (t *Tree) Delete(x int) {
	t.root = del(t.root, x)
}

func del(node *Node, x int) *Node {

	if node == nil {
		return nil
	} else if node.value > x {
		// Find in the left branch
		node.left = del(node.left, x)
		return node
	} else if node.value < x {
		// Find in the right branch
		node.right = del(node.right, x)
		return node
	} else if node.value == x {
		// If the node has no child then remove it.
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			// Have two child
			// If the current node has two children, we delete
			// the current node by calling the lift function(below)
			// which changes the current node's value to the value
			// of its successor node:
			//node.right = lift(node.right, node)
		}
	}
	return nil
}

func lift(node, nodeToDelete *Node) *Node {
	if node.left != nil {
		node.left = lift(node.left, nodeToDelete)
		return node
	} else {
		nodeToDelete.value = node.value
		return node.right
	}
}


// Using recursion
// Step 1: Traverse the tree and inspect the value of the node.
// Step 2: If the node is nil or equal to the data to be search return node
// Step 3: If the node's value is greater than x value then traverse the left
//				 branch.
// Step 4: If the node's value is lesser than x value then traverse the right
//         branch.
func search(node *Node, x int) *Node {
	if node == nil || node.value == x {
		return node
	} else if node.value > x {
		// Traverse in the left branch
		return search(node.left, x)
	} else {
		// Traverse in the right branch
		return search(node.right, x)
	}
}

// Step 1: Check the value of the node to the inserting value
//	Cond 1: If the node value is lesser than the inserting value
//		Cond 1: Check if the left node of the current node is empty
//						then insert the new node to the left node
//    Cond 2: If not empty, then recursively call the insert with the current left node
//  Cond 2: if the node is greater than the inserting node
//		Cond 1: Check if the right node of the current node is empty
//						then insert the new node to the right node
//    Cond 2: If not empty, then recursively call the insert with the current right node
func insert(n *Node, data int) {
	if n.value > data {
		if n.left != nil {
			insert(n.left, data)
		} else {
			n.left = &Node{
				value: data,
			}
		}
	} else {
		if n.right != nil {
			insert(n.right, data)
		} else {
			n.right = &Node{
				value: data,
			}
		}
	}
}

func (t *Tree) isEmpty() bool {
	return t.size == 0
}
func (t *Tree) Size() int {
	return t.size
}
