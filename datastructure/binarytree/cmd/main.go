package main

import (
	"fmt"
	"github.com/jayvib/ds-algo/datastructure/binarytree"
)

func main() {
	treeNode := &binarytree.Node{Value: 10}

	binarytree.Insert(5, treeNode)
	fmt.Println(treeNode.LeftChild.Value)


	node := binarytree.Search(5, treeNode)

	fmt.Println(node.Value)
}
