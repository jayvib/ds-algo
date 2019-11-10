package main

import (
	"fmt"
	"github.com/jayvib/ds-algo/datastructure/binarytree/binarysearchtree"
)

func main() {
	t := binarysearchtree.New()
	t.InsertElement(1, 8)
	t.InsertElement(2, 3)
	t.InsertElement(3, 4)
	t.InsertElement(4, 10)

	t.InOrderTraverseTree(func(n *binarysearchtree.TreeNode){
		fmt.Println(n.Value())
	})
}
