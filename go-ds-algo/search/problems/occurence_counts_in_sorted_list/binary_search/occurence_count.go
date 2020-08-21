package binary_search

func OccurenceCount(data []int, x int) int {
	size := len(data)
	// Find the first index occurrence of x
	fi := findFirstIndex(data, 0, size-1, x)

	// Find the last index occurrence of x
	li := findLastIndex(data, 0, size-1, x)

	// Subtract: last-first+1
	return li - fi + 1
}

func findFirstIndex(data []int, low, high int, key int) int {
	if low >= high {
		return -1
	}

	mid := low + (high-low)/2
	if data[mid] == key && (low == mid || data[mid-1] != key) {
		return mid
	}

	if data[mid] >= key {
		// Find in the lower bound
		return findFirstIndex(data, low, mid-1, key)
	}
	return findFirstIndex(data, mid+1, high, key)
}

func findLastIndex(data []int, low, high int, key int) int {
	if low >= high {
		return -1
	}

	mid := low + (high-low)/2

	if data[mid] == key && (mid == high || data[mid+1] != key) {
		return mid
	}

	if data[mid] >= key {
		return findLastIndex(data, mid+1, high, key)
	}
	return findLastIndex(data, low, mid-1, key)
}
