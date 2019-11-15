package main

import (
	"fmt"

	"github.com/jayvib/ds-algo/datastructure/binarytree/binarytree.v2"
)

func main() {
	t := binarytree.New()
	t.InsertElement(1, 1)
	t.InsertElement(3, 3)
	t.InsertElement(5, 5)
	t.InsertElement(2, 2)
	t.InsertElement(4, 4)

	fmt.Println("-----IN ORDER TRAVERSE-----")
	t.InOrderTraverseTree(func(n *binarytree.TreeNode) {
		fmt.Println(n.Value())
	})

	fmt.Println("-----PRE ORDER TRAVERSE-----")
	t.PreOrderTraverseTree(func(n *binarytree.TreeNode) {
		fmt.Println(n.Value())
	})

	fmt.Println("-----POST ORDER TRAVERSE-----")
	t.PostOrderTraverseTree(func(n *binarytree.TreeNode) {
		fmt.Println(n.Value())
	})

	fmt.Println("Minimum Value:", *t.MinNode())

	fmt.Println("Maximum Value:", *t.MaxNode())

	fmt.Println("Found 3?", t.SearchNode(3))
	fmt.Println("Found 6?", t.SearchNode(6))
}
