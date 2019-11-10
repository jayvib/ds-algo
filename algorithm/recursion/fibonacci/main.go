package main

import "fmt"

func fibonacci(k int) int {
	if k <= 1 {
		return 1
	}
	return fibonacci(k-1) + fibonacci(k-2)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ",fibonacci(i))
	}

}