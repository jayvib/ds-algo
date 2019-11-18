package main

import (
	"container/list"
	"fmt"
)

type List struct {
	l *list.List
}

func (l *List) Print() {
	for n := l.l.Front(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}
}

func main() {
	l := list.New()
	l.PushBack(3)
	l.PushBack(2)
	l.PushBack(1)
	myList := &List{l:l}
	myList.Print()
}
