package main

import (
	"fmt"
	"github.com/jayvib/ds-algo/datastructure/linkedlist"
	"github.com/jayvib/golog"
)

func main() {
	golog.SetLevel(golog.TraceLevel)
	l := linkedlist.New()
	l.PushFront(0)
	l.PushFront(2)
	l.PushFront(3)
	l.PushBack(4)

	l.Print()

	fmt.Println("is 2?", l.Get(1)==2)
	fmt.Println("is found 2?", l.IsExists(2))
	fmt.Println(l.Find(3))
}
