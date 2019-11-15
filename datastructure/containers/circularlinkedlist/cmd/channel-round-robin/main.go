package main

import (
	"container/ring"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chanList := ring.New(3)
	chanPool := []chan int {
		make(chan int),
		make(chan int),
		make(chan int),
	}


	for i := 0; i < 3; i++ {
		chanList.Value = chanPool[i]
		chanList = chanList.Next()
	}

	for i := 0; i < 3; i++ {
		go func(c int) {
			for v := range chanPool[c] {
				fmt.Printf("%d: %v\n", c, v)
			}
			fmt.Println("exiting:", c)
		}(i)
	}

	count := 0
	for r := chanList.Next();; r = r.Next() {
		if count == 20 {
			break
		}
		count++
		ch := r.Value.(chan int)
		ch <- rand.Int()
		time.Sleep(100 * time.Millisecond)
	}

	// closing the channels
	for _, ch := range chanPool {
		close(ch)
	}
	time.Sleep(3 * time.Second)
}
