package main

import (
	"fmt"
	"time"
)

func countdown(sec int) {
	// Base case
	if sec < 0 {
		return
	}
	time.Sleep(time.Second)
	fmt.Println(sec)
	countdown(sec-1)
}

func main() {
	countdown(10)
}
