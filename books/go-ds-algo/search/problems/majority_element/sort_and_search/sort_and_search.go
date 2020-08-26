package sort_and_search

import "sort"

func FindMajority(ints []int) (element int, found bool) {
	sort.Ints(ints)
	size := len(ints)
	majorityIndex := ints[size/2]
	candidate := ints[majorityIndex]
	count := 0
	for i := 0; i < size; i++ {
		if ints[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		element = candidate
		found = true
	}
	return
}
