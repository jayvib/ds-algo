package main

import (
	"fmt"
	"github.com/jayvib/ds-algo/datastructure/linkedlist/review"
)

func main() {
	l := &review.List{}
	l.AddToHead(1)
	l.AddToHead(1)
	l.AddToHead(1)
	l.AddToHead(2)
	l.AddToHead(3)
	l.AddToHead(4)
	fmt.Println(l.GetHead())
	l.Iterate()
	fmt.Println(l.LastNode())
}
