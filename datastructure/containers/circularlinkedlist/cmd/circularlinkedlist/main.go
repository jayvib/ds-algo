package main

import (
	"container/ring"
	"fmt"
)

func main() {
	ExampleRingMove()
}

func ExampleRingMove() {
	size := 10
	circularList := ring.New(size)

	for i := 0; i < size; i++ {
		circularList.Value = i
		circularList = circularList.Next()
	}

	circularList.Do(func(v interface{}){
		fmt.Println(v)
	})

	circularList = circularList.Move(2)
	fmt.Println("-------------")
	circularList.Do(func(v interface{}){
		fmt.Println(v)
	})
}

func ExampleRingDo() {
	circularList := ring.New(5)

	n := circularList.Len()
	for i := 0; i < n; i++ {
		circularList.Value = i
		circularList = circularList.Next()
	}

	circularList.Do(func(p interface{}){
		fmt.Println(p)
	})

}

