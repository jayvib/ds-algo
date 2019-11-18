package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
	"log"
	"sync"
	"time"
)

func queueExample() {
	q := queue.New(100)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		count := 0
		for i := 0; i < 100; i++ {
			count++
			if (count % 5) == 0 {
				time.Sleep(100 * time.Millisecond)
			}
			err := q.Put(i)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	for {
		items, err := q.Poll(5, 3 * time.Second)
		if err != nil {
			if err == queue.ErrEmptyQueue {
				fmt.Println("empty queue")
				break
			}
			break
		}

		fmt.Printf("%#v\n", items)
	}

	disposed := q.Dispose()
	fmt.Printf("%#v", disposed)
	wg.Wait()

}

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
