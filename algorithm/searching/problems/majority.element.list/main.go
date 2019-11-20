package main

import "sort"

// Majority element in a list
// Given a list of n elements. Find the majority element, which appears more than n/2
// times. Return 0 in case there is no majority element.

func FindMajorityElementBruteForce(elements []int) (majority int, exists bool) {
	max := elements[0]
	maxCount := 1

	size := len(elements)
	for i := 0; i < size; i++ {
		currentCount := 1
		for j := i+1; j < size;j++ {
			if elements[i] == elements[j] {
				currentCount++
			}
		}

		if currentCount > maxCount {
			max = elements[i]
			maxCount = currentCount
		}
	}

	if maxCount > (size /2)	 {
		return max, true
	}

	return
}

func FindMajorityElementSorted(elements []int) (majority int, exists bool) {
	if !sort.IsSorted(sort.IntSlice(elements))	 {
		sort.Ints(elements)
	}

	mid := len(elements)/2
	candidate := elements[mid]
	count := 1
	for i := 0; i < len(elements); i++ {
		if elements[i] == candidate {
			count++
		}
	}
	if count > len(elements)/2 {
		return candidate, true
	}
	return
}

func FindMajorityElementSortedReview(elements []int) (int, bool) {
	if !sort.IsSorted(sort.IntSlice(elements)) {
		sort.Ints(elements)
	}
	size := len(elements)
	majorityElement := size/2
	candidate := elements[majorityElement]
	count := 1
	for i := 0; i < size; i++ {
		if elements[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		return candidate, true
	}
	return 0, false
}