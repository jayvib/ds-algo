package bst

import "fmt"

type TreeNode struct {
	key       int
	value     int
	leftTree  *TreeNode
	rightTree *TreeNode
}

func (n *TreeNode) Key() int {
	return n.key
}

func (n *TreeNode) Value() int {
	return n.value
}

func (n *TreeNode) String() string {
	return fmt.Sprintf("key: %d value: %d", n.key, n.value)
}

type BST struct {
	rootNode *TreeNode
}

func (b *BST) Insert(key, value int) {
	newTreeNode := &TreeNode{key: key, value: value}
	if b.rootNode == nil {
		b.rootNode = newTreeNode
	} else {
		insertTreeNode(b.rootNode, newTreeNode)
	}
}

func (b *BST) InOrderTraverse(fn func(n *TreeNode)) {
	inOrderTraverse(b.rootNode, fn)
}

func (b *BST) PreOrderTraverse(fn func(n *TreeNode)) {
	preOrderTraverse(b.rootNode, fn)
}

func (b *BST) PostOrderTraverse(fn func(n *TreeNode)) {
	postOrderTraverse(b.rootNode, fn)
}

func (b *BST) MinNode() *TreeNode {

	// find the minimum value which is the left most of the
	// subtree
	currentNode := b.rootNode
	for {
		if currentNode.leftTree == nil {
			return currentNode
		}

		currentNode = currentNode.leftTree
	}
}

func (b *BST) MaxNode() *TreeNode {
	if b.rootNode == nil {
		return nil
	}

	currentNode := b.rootNode
	for {
		if currentNode.rightTree == nil {
			return currentNode
		}
		currentNode = currentNode.rightTree
	}
}

func (b *BST) SearchNode(key int) *TreeNode {
	return searchNode(b.rootNode, key)
}

func searchNode(n *TreeNode, key int) *TreeNode {
	switch {
	case n == nil:
		return n
	case n.key > key:
		// find in the left
		return searchNode(n.leftTree, key)
	case n.key < key:
		return searchNode(n.rightTree, key)
	default:
		return n
	}
}

func postOrderTraverse(n *TreeNode, fn func(n *TreeNode)) {
	if n != nil {
		postOrderTraverse(n.leftTree, fn)
		postOrderTraverse(n.rightTree, fn)
		fn(n)
	}
}

func preOrderTraverse(n *TreeNode, fn func(n *TreeNode)) {
	if n != nil {
		fn(n)
		preOrderTraverse(n.leftTree, fn)
		preOrderTraverse(n.rightTree, fn)
	}
}

func inOrderTraverse(node *TreeNode,fn func(n *TreeNode)) {
	if node != nil {
		inOrderTraverse(node.leftTree, fn)
		fn(node)
		inOrderTraverse(node.rightTree, fn)
	}
}

func insertTreeNode(root, newNode *TreeNode) {
	if newNode.key < root.key {
		// Put the new node in the right subtree
		if root.leftTree == nil {
			root.leftTree = newNode
		} else {
			insertTreeNode(root.leftTree, newNode)
		}
	} else {
		if root.rightTree == nil {
			root.rightTree = newNode
		}	else {
			insertTreeNode(root.rightTree, newNode)
		}
	}
}