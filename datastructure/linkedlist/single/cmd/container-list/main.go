package main

import (
	"container/list"
	"fmt"
)

func main() {
	intList := list.New()
	intList.PushBack(1)
	intList.PushBack(2)
	intList.PushBack(3)

	for next := intList.Front(); next != nil; next = next.Next() {
		fmt.Println(next.Value)
	}
}
