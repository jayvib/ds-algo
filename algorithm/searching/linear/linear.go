package linear

import "sort"

// Time Complexity: O(n)
func Search(items []int, find int) bool {
	for i := range items {
		if items[i] == find {
			return true
		}
	}
	return false
}

// Time Complexity: O(n) as the worst case is when the item is in the last index
// Space Complexity: O(1) since no memory is needed to allocate the list
func SearchSorted(items []int, find int) bool {
	if !sort.IsSorted(sort.IntSlice(items)) {
		sort.Ints(items)
	}

	for _, i := range items {
		if i == find {
			return true
		}
		if i > find {
			return false
		}
	}
	return false
}