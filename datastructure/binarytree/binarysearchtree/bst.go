package binarysearchtree

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
	lock     sync.RWMutex
}

func (t *BinarySearchTree) InsertElement(key int, value int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	node := &TreeNode{key: key, value: value}
	if t.rootNode == nil {
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

func (t *BinarySearchTree) InOrderTraverseTree(fn func(n *TreeNode)) {
	t.lock.Lock()
	defer t.lock.Unlock()
	inOrderTraverseTree(t.rootNode, fn)
}

func inOrderTraverseTree(node *TreeNode, fn func(n *TreeNode)) {
	if node != nil {
		inOrderTraverseTree(node.leftNode, fn)
		fn(node)
		inOrderTraverseTree(node.rightNode, fn)
	}
}
