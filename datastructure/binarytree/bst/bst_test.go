package bst

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	tree := &BST{}
	tree.Insert(2, 5)
	tree.Insert(5, 8)
	tree.Insert(4, 7)
	tree.Insert(3, 6)

	tree.InOrderTraverse(func(n *TreeNode){
		fmt.Print(n.value, " ")
	})
	fmt.Println()
	tree.PreOrderTraverse(func(n *TreeNode){
		fmt.Print(n.value, " ")
	})
	fmt.Println()
	tree.PostOrderTraverse(func(n *TreeNode){
		fmt.Print(n.value, " ")
	})
	fmt.Println()

	fmt.Println("Minimum:", tree.MinNode())
	fmt.Println("Maximum:", tree.MaxNode())
	fmt.Println("Searching for 4:", tree.SearchNode(10))
}
