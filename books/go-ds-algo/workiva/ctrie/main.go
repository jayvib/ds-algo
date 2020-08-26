package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/trie/ctrie"
	"strconv"
	"time"
)

func main() {
	c := ctrie.New(nil)
	go func() {
		for i := 0; i < 100; i++ {
			ns := strconv.Itoa(i)
			c.Insert([]byte(ns), ns)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	time.Sleep(3 * time.Second)
	doneChan := make(chan struct{})
	defer close(doneChan)
	iter := c.Iterator(doneChan)
	for v := range iter {
		fmt.Println(v.Value)
	}
}
