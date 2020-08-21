package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/tree/avl"
	"strings"
)

type entry struct {
	key string
	value string
}

func (e *entry) Compare(from avl.Entry) int {
	incomingEntry := from.(*entry)
	return strings.Compare(e.key, incomingEntry.key)
}

func main() {
	tree := avl.NewImmutable()

	item1 := &entry{key:"a", value:"Hello"}
	newTree, _ := tree.Insert(item1)
	fmt.Println(newTree.Len())
	fmt.Println(tree.Len())
	newTree.Get()
}
