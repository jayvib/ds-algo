package binary

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
