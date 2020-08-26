package brute_force

import "sort"

func FindMajority(ints []int) (int, bool) {
	size := len(ints)
	// Find the max recurring element
	elem, count := findMaxRecurringElement(ints)
	if count > size/2 {
		return elem, true
	}
	return 0, false
}

func findMaxRecurringElement(ints []int) (element, count int) {
	// Gonna use sort and search

	// Is ints sorted?
	if !sort.IsSorted(sort.IntSlice(ints)) {
		sort.Ints(ints)
	}

	max := ints[0]
	maxCount := 1

	current := ints[0]
	currentCount := 1

	size := len(ints)
	for i := 1; i < size; i++ {
		if ints[i] == ints[i-1] {
			currentCount++
		}	else {
			current = ints[i]
			currentCount = 1
		}
		if currentCount > maxCount {
			max = current
			maxCount = currentCount
		}
	}
	return max, maxCount
}
