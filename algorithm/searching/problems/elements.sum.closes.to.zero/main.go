package main

import (
	"fmt"
	"math"
)

// Problem: Given a list of integers, both +ve and -ve.
// You need to find the two elements such that their sum is closest to zero.

func MinAbsSumPairBrute(data []int) {
	size := len(data)
	if size < 2 {
		fmt.Println("Invalid Input")
		return
	}

	// Initial value
	minFirst, minSecond := 0, 1
	minSum := abs(data[minFirst] + data[minSecond])

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			sum := data[i] + data[j]
			if sum < minSum {
				minSum = sum
				minFirst = i
				minSecond = j
			}
		}
	}
	fmt.Printf("The two elements with minimum sum: %d, %d\n", data[minFirst], data[minSecond])
}

func MinAbsSumPairSorted(data []int) {
	size := len(data)
	if size < 2 {
		fmt.Println("invalid input")
		return
	}

	minFirst, minSecond := 0, size-1
	minSum := abs(data[minFirst] + data[minSecond])

	for l, r := 0, size-1; l < r; {
		sum := abs(data[l] + data[r])
		if sum < minSum {
			minSum = sum
			minFirst, minSecond = l, r
		}

		if sum < 0 {
			l++
		} else if sum > 0 {
			r--
		} else {
			break
		}
	}
	fmt.Printf("The two elements with minimum sum: %d, %d\n", data[minFirst], data[minSecond])
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	input := []int{2, 5, 8, 9, 1}
	MinAbsSumPairSorted(input)
	MinAbsSumPairBrute(input)
}
