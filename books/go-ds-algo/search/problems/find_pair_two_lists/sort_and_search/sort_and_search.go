package sort_and_search

import (
	"sort"
)

func FindPair(setA, setB []int, value int) (pairs [][2]int, ok bool) {
	sort.Ints(setB) // O(logn)
	for i := 0; i < len(setA); i++ {
		x := setA[i]
		y := value- x
		if isFound := search(setB, y); isFound { // O(nlogn)
			pairs = append(pairs, [2]int{x, y})
		}
	}
	return pairs, len(pairs)>0
}

func search(ints []int, n int) (ok bool) {
	size := len(ints)
	low := 0
	high := size-1
	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if ints[mid] == n {
			return true
		} else if ints[mid] < n {
			low = mid+1
		} else {
			high = mid-1
		}
	}
	return false
}

