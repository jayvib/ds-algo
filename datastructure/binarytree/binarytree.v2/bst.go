package binarytree

import "sync"

type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

func (t *TreeNode) Key() int {
	return t.key
}

func (t *TreeNode) Value() int {
	return t.value
}

func New() *BinarySearchTree {
	return &BinarySearchTree{}
}

type BinarySearchTree struct {
	rootNode *TreeNode
	rwmu     sync.RWMutex
}

func (t *BinarySearchTree) InsertElement(key int, value int) {
	t.rwmu.Lock()
	defer t.rwmu.Unlock()
	node := &TreeNode{key: key, value: value}
	if t.rootNode == nil { // meaning this will be the root node.
		t.rootNode = node
	} else {
		insertTreeNode(t.rootNode, node)
	}
}

func insertTreeNode(rootNode *TreeNode, newNode *TreeNode) {
	if newNode.key < rootNode.key { // so the new node is for left node.
		if rootNode.leftNode == nil {
			rootNode.leftNode = newNode
		} else {
			// find the left location to the next level
			insertTreeNode(rootNode.leftNode, newNode)
		}

		// so the new node is for right branch.
	} else {

		// This will be the base case for the right branch.
		// When the root's node right is empty.
		if rootNode.rightNode == nil {
			rootNode.rightNode = newNode

			// Then keep find the correct slot for the new node
		} else {
			insertTreeNode(rootNode.rightNode, newNode)
		}
	}
}

func insertTreeNode2(rootNode *TreeNode, newNode *TreeNode) {
	// if the new nodes key is less then the current node.
	// then put the new node to the left side.
	if newNode.key < rootNode.key {
		// if the current root node's left node is empty
		// then the new node will become the new left node of the
		// current tree.
		// This is the base case for the left side of the node.
		if rootNode.leftNode == nil {
			rootNode.leftNode = newNode
		} else {
			// Find slot in the next level
			insertTreeNode2(rootNode.leftNode, newNode)
		}

		// so the new node key is greater than the current node key value
		// then assign it to the right side of the node.
	} else {

		// if the current node's right side node is empty.
		// then let the new new node the new right side node.
		// This is the base case for the right side of the node.
		if rootNode.rightNode == nil {
			rootNode.rightNode = newNode
		} else {
			insertTreeNode2(rootNode.rightNode, newNode)
		}
	}
}

func (t *BinarySearchTree) InOrderTraverseTree(fn func(n *TreeNode)) {
	t.rwmu.Lock()
	defer t.rwmu.Unlock()
	inOrderTraverseTree(t.rootNode, fn)
}

func inOrderTraverseTree(node *TreeNode, fn func(n *TreeNode)) {
	if node != nil {
		inOrderTraverseTree(node.leftNode, fn)
		fn(node)
		inOrderTraverseTree(node.rightNode, fn)
	}
}

func (t *BinarySearchTree) PreOrderTraverseTree(fn func(n *TreeNode)) {
	t.rwmu.Lock()
	defer t.rwmu.Unlock()
	preOrderTraverseTree(t.rootNode, fn)
}

func preOrderTraverseTree(rootNode *TreeNode, fn func(n *TreeNode)) {
	if rootNode != nil {
		fn(rootNode)
		preOrderTraverseTree(rootNode.leftNode, fn)
		preOrderTraverseTree(rootNode.rightNode, fn)
	}
}

func (t *BinarySearchTree) PostOrderTraverseTree(fn func(n *TreeNode)) {
	t.rwmu.Lock()
	defer t.rwmu.Unlock()
	postOrderTraversTree(t.rootNode, fn)
}

func postOrderTraversTree(rootNode *TreeNode, fn func(n *TreeNode)) {
	if rootNode != nil {
		postOrderTraversTree(rootNode.leftNode, fn)
		postOrderTraversTree(rootNode.rightNode, fn)
		fn(rootNode)
	}
}

func (t *BinarySearchTree) MinNode() *int {
	t.rwmu.RLock()
	defer t.rwmu.RUnlock()
	currNode := t.rootNode
	if currNode == nil {
		return nil
	}

	for {
		if currNode.leftNode == nil {
			return &currNode.value
		}
		currNode = currNode.leftNode
	}
}

func (t *BinarySearchTree) MaxNode() *int {
	t.rwmu.RLock()
	defer t.rwmu.RUnlock()
	currNode := t.rootNode
	if currNode == nil {
		return nil
	}

	for {
		if currNode.rightNode == nil {
			return &currNode.value
		}
		currNode = currNode.rightNode
	}
}

func (t *BinarySearchTree) SearchNode(key int) bool {
	t.rwmu.RLock()
	defer t.rwmu.RUnlock()
	return searchNode(t.rootNode, key)
}

func searchNode(treeNode *TreeNode, key int) bool {
	switch {
	case treeNode == nil:
		return false
	case key < treeNode.key:
		return searchNode(treeNode.leftNode, key)
	case key > treeNode.key:
		return searchNode(treeNode.rightNode, key)
	default:
		return true
	}
}

func deleteNode(treeNode *TreeNode, key int) *TreeNode {
	if treeNode == nil {
		return nil
	}

	// Rule 1: If the node being deleted has one child, delete it and plug the
	// child into the spot where the deleted node was.
	if key < treeNode.key {
		treeNode.leftNode = deleteNode(treeNode.leftNode, key)
		return treeNode
	}

	// Rule 2: If the node being deleted has one child, delete it and plug the
	// child into the spot where the deleted node was.
	if key > treeNode.key {
		treeNode.rightNode = deleteNode(treeNode.rightNode, key)
		return treeNode
	}
	// ########################################################
	// ##### We are in the current node that will be delete ###
	// ########################################################

	// Rule 3: If the node being deleted has no children, simply delete it.
	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		treeNode = nil
		return nil
	}

	// when the current node's left is empty.
	// the right node will become this current node.
	if treeNode.leftNode == nil {
		treeNode = treeNode.rightNode
		return treeNode
	}

	// when the current node's right is empty.
	// the left node will become this current node.
	if treeNode.rightNode == nil {
		treeNode = treeNode.leftNode
		return treeNode
	}

	// When the current tree node to be delete
	// has 2 children.
	//
	// Rule 4: When deleting a node with two children, replace the
	// deleted node with the successor node. The successor node
	// is the child node whose value is the least of all values that
	// are greater than the deleted node.

	leftMostRightNode := treeNode.rightNode // get the right node.

	// find the smallest value on the right side
	for {
		if leftMostRightNode != nil && leftMostRightNode.leftNode != nil {
			leftMostRightNode = leftMostRightNode.leftNode
		} else {
			break
		}
	}

	// so the right node's smallest value will become the value for this node.
	treeNode.key, treeNode.value = leftMostRightNode.key, leftMostRightNode.value

	// Rule: If the successor node has a right child, after plugging the successor node into
	// the post of the deleted node, take the right child of the successor node and turn it
	// into the left child of the parent of the successor node.
	treeNode.rightNode = deleteNode(treeNode.rightNode, treeNode.key)
	return treeNode
}
