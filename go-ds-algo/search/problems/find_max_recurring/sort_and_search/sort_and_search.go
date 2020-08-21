package sort_and_search

import "sort"

func FindMaxRecurring(ints []int) int {
	// Is sorted?
	if !sort.IsSorted(sort.IntSlice(ints)) {
		sort.Ints(ints)
	}
	return findMaxRecurringSorted(ints)
}

// findMaxRecurringSorted finds the max existing
// item in ints.
//
// ints must be sorted already.
func findMaxRecurringSorted(ints []int) int {
	size := len(ints)
	max := ints[0] // assume first value as the max
	maxCount := 1

	current := ints[0] // pre-selection
	currentCount := 1

	for i := 1; i < size; i++ {
		if ints[i] == ints[i-1] {
			currentCount++
		} else {
			currentCount = 1
			current = ints[i]
		}
		// Check if this current count is
		// greater with the existing max
		// count
		if currentCount > maxCount {
			max = current
			maxCount = currentCount
		}
	}
	return max
}
