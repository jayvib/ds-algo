package binary

func Search(data []int, x int) (index int) {
	return binarySearchIteration(data, 0, len(data)-1, x)
}

func binarySearchIteration(data []int, low, high, x int) (index int) {

	for low <= high {
		mid := (high + low) / 2

		if data[mid] == x {
			return mid
		}

		if data[mid] > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func binarySearchRecursion(data []int, low, high, x int) (index int) {

	for low <= high {
		mid := low + (high-low)/2

		if data[mid] == x {
			return mid
		}

		if data[mid] > x {
			// search in the lower half
			return binarySearchRecursion(data, low, mid-1, x)
		} else {
			return binarySearchRecursion(data, mid+1, high, x)
		}

	}
	return -1
}
