package binary

import "sort"

// Time Complexity - O(logn)

// Binary search requires the list to be sorted otherwise
// binary search cannot be applied.
func Search(array []int, find int) bool {
	return binarySearch(array, 0, len(array)-1, find)
}

func binarySearch(data []int, low, high, value int) bool {
	// base case is when the low is greater than the high
	if low > high {
		return false
	}

	mid := low + (high-low)/2 // calculate the middle index

	if data[mid] == value {
		return true

	// find in the upper partition
	} else if data[mid] < value {
		return binarySearch(data, mid+1, high, value)
	} else if data[mid] > value {
		return binarySearch(data, low, mid-1, value)
	}
	return false
}

func SearchIterative(array []int, find int) bool {
	var low, mid, high int
	low, mid, high = 0, 0, len(array)-1

	for low <= high {
		// calculate the mid
		mid = low + (high-low)/2
		if array[mid] == find {
			return true
		} else if array[mid] > find {
			// find in the lower partition
			high = mid-1
		} else if array[mid] < find {
			// find in the upper partition
			low = mid+1
		}
	}
	return false
}

func SearchReview(array []int, find int) bool {
	if !sort.IsSorted(sort.IntSlice(array)) {
		sort.Ints(array)
	}
	return binarySearchReview(array, 0, len(array)-1, find)
}

func binarySearchReview(array []int, low, high, value int) bool {
	if low > high { // this will be the base case
		return false
	}

	// calculate the mid index
	mid := low + (high-low)/2

	if array[mid] == value {
		return true
	}

	if array[mid] > value {
		// find in lower partition
		return binarySearchReview(array, low, mid-1, value)
	} else if array[mid] < value { // find in upper partition
		return binarySearchReview(array, mid+1, high, value)
	}
	return false
}

