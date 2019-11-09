package main

import (
	"fmt"
	"sync"
	"time"
)

func Countdown(sec int) {
	ticker := time.NewTicker(time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		count := sec
		for {
			select {
			case <-ticker.C:
				fmt.Println(count)

				if count == 0 {
					return
				}
			}
		}
	}()
	wg.Wait()
}

func CountdownRecursion(number int) {
	if number < 0 {
		return
	}
	fmt.Println(number)
	CountdownRecursion(number-1)
}

func Factorial(number int) int {
	if number == 1 {
		return 1
	}

	return number * Factorial(number-1)
}

func main() {
	fmt.Println("Factorial of 5 is", Factorial(5))
}
