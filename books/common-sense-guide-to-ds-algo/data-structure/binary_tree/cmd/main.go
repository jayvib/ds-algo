package main

import (
	"fmt"
	"github.com/jayvib/ds-algo/books/common-sense-guide-to-ds-algo/data-structure/binary_tree"
)

func traverseAndPrint(node *binary_tree.Node) {
	if node == nil {
		return
	}

	traverseAndPrint(node.Left())
	fmt.Println(node.Value())
	traverseAndPrint(node.Right())
}

func preOrderTraverseAndPrint(node *binary_tree.Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Value())
	traverseAndPrint(node.Left())
	traverseAndPrint(node.Right())
}

func postOrderTraverseAndPrint(node *binary_tree.Node) {
	if node == nil {
		return
	}
	traverseAndPrint(node.Left())
	traverseAndPrint(node.Right())
	fmt.Println(node.Value())
}

func main() {
	t := binary_tree.Tree{}
	t.Insert(5)
	t.Insert(4)
	t.Insert(6)
	t.Insert(3)
	t.Insert(2)

	//traverseAndPrint(t.Root())
	preOrderTraverseAndPrint(t.Root())
}
