package binarytree

type Tree struct {
	root *Node
}

type Node struct {
	Value      int
	LeftChild  *Node
	RightChild *Node
}

func Insert(value int, node *Node) {
	if value < node.Value {
		// Going to insert it on the left side
		if node.LeftChild == nil {
			// create a new node
			node.LeftChild = &Node{Value: value}
		} else {
			Insert(value, node.LeftChild)
		}
	} else if value > node.Value {
		if node.RightChild == nil {
			node.RightChild = &Node{Value: value}
		} else {
			Insert(value, node.RightChild)
		}
	} // what if equal value?
}

func Search(value int, node *Node) *Node {
	// Base Case
	if value == node.Value {
		return  node
	}

	// if the Value is less than the current node Value
	// then search in the left branch.
	if value < node.Value {
		return Search(value, node.LeftChild)
	} else {
		return Search(value, node.RightChild)
	}

	return nil
}

func Delete(value int, node *Node) *Node {
	// The base case is when we've hit the bottom of the tree,
	// and the parent node has no children.
	if node == nil {
		return nil
	}

	// When the value to delete is less than the current node value
	// then find the value in the left side of the node then return
	// the node.
	if value < node.Value {
		node.LeftChild = Delete(value, node)
		return node

	// When the value to delete is greater then the current node value
	// then find the value in the right side of the node then return
	// the node.
	} else if value > node.Value {
		node.RightChild = Delete(value, node.RightChild)
		return node

	// When the value to delete is equal to the current node.
	} else {

		// Check if the current node's left child is empty.
		if node.LeftChild == nil {
			// return the right child
			return node.RightChild

		// Check if the current node's right child is empty
		} else if node.RightChild == nil {
			// return the left child
			return node.LeftChild

		// so the current node has 2 children.
		} else {
			node.RightChild = lift(node.RightChild, node)
			return node
		}
	}
	return nil
}

func lift(node, nodeToDelete *Node) *Node {
	if node.LeftChild != nil {
		node.LeftChild = lift(node.LeftChild, nodeToDelete)
		return node
	}	else {
		nodeToDelete.Value = node.Value
		return node.RightChild
	}
}