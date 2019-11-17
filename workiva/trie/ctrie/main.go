package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/trie/ctrie"
)

// When to use Trie Data Structure?

func main() {
	ct := ctrie.New(nil)
	//ct.Insert([]byte("luffy"), "monkey")
	//ct.Insert([]byte("sanji"), "vinsmoke")
	//ct.Insert([]byte("zoro"), "roronoa")
	ct.Insert([]byte("a"), "a")
	ct.Insert([]byte("b"), "b")
	ct.Insert([]byte("c"), "c")
	ct.Insert([]byte("d"), "d")
	ct.Remove([]byte("c"))
	for v := range ct.Iterator(nil) {
		fmt.Println(string(v.Key), v.Value)
	}
}
