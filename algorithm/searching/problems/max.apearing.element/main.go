package main

import "sort"

// Problem:
// Given a list of n numbers, find the element, which appears maximum number of times.


// Steps:
// Get initial maximum number
// iterate the list using two for loops.
// Time Complexity:  O(n^2)
// Space Complexity: O(n^2)
func GetMaxAppearingElementBruteForce(elements []int) int {
	// Set initial max value
	max := elements[0]
	count := 1
	maxCount := 1

	for i := 0; i < len(elements); i++ {
		count = 1
		for j := i+1; j < len(elements); j++ {
			if elements[i] == elements[j] {
				count++
			}
		}

		if count > maxCount {
			maxCount = count
			max = elements[i]
		}
	}

	return max
}

func GetMaxAppearingElementSorted(elements []int) int {
	if !sort.IsSorted(sort.IntSlice(elements)) {
		sort.Ints(elements)
	}
	max := elements[0]
	maxCount := 1
	current := elements[0]
	currentCount := 1
	for i := 1; i < len(elements);i++ {
		if elements[i] == elements[i-1] {
			currentCount++
		} else {
			currentCount = 1
			current = elements[i]
		}
		if currentCount > maxCount {
			 maxCount = currentCount
			 max = current
		}
	}
	return max
}

func main() {

}

func GetMaxAppearingElementBruteForceReview(elements []int) int {
	// set initial maximum number
	max := elements[0]
	maxCount := 1

	for i := 0; i < len(elements); i++ {
		count := 1
		for j := i+1; j < len(elements); j++ {
			if elements[i] == elements[j] {
				count++
			}
		}

		// compare the current max count and the current value
		// frequency
		if count > maxCount {
			max = elements[i]
			maxCount = count
		}
	}
	return max
}

func GetAppearingElementSortedReview(elements []int) int {
	current := elements[0]
	currentCount := 1
	max := elements[0]
	maxCount := 1
	for i := 1; i < len(elements);i++ {
		if elements[i] == elements[i-1] {
			currentCount++
		} else {
			current = elements[i]
			currentCount = 1
		}

		if currentCount > maxCount {
			max = current
			maxCount = currentCount
		}
	}
	return max
}
