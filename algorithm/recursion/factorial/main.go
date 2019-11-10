package main

import "fmt"

func Factorial(number int) int {
	if number == 1 {
		return 1
	}

	return number * Factorial(number-1)
}

func main() {
	fmt.Println("Factorial of 5 is", Factorial(5))
}
