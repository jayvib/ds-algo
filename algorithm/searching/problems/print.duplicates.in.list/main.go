package main

import (
	"fmt"
	"sort"
)

// Print duplicates in the list
const (
	brute = iota
	sorted
	hash
)
func PrintDuplicates(array []int, approach int) {
	switch approach {
	case brute:
		printDuplicatesBruteForce(array)
	case sorted:
		printDuplicateSorted(array)
	case hash:
		printDuplicateHash(array)
	}
}

func printDuplicatesBruteForce(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i+1; j < len(array); j++ {
			if array[i] == array[j] {
				fmt.Printf("%d ", array[i])
				// no need to continue..... move to the next i items.
			}
		}
	}
}

func printDuplicateSorted(array []int) {
	sort.Ints(array)
	for i := 1; i < len(array); i++ {
		if array[i] == array[i-1] {
			fmt.Printf("%d ", array[i])
		}
	}
}

// Time Complexity: 0(n)
// Space Complexity: O(n)
func printDuplicateHash(array []int) {
	m := make(map[int]struct{})
	for _, a := range array {
		_, ok := m[a]
		if ok {
			fmt.Printf("%d ", a)
			continue
		}
		m[a] = struct{}{}
	}
}

func main() {
	nums := []int{1, 3, 4, 3, 4, 2, 10, 8, 8, 10, 3}
	PrintDuplicates(nums, hash)
}
