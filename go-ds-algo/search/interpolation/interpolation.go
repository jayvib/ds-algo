package interpolation

func Search(data []int, value int) (index int, found bool) {
	size := len(data)
	low := 0
	high := size-1
	mid := -1
	index = -1
	for low <= high && data[high] >= value && data[low] <= value {
		// Start searching data from middle of the list.
		mid = low + ((high-low)/(data[high]-data[low])) * (value - data[low])
		probableValue := data[mid]
		if probableValue == value {
			return mid, true
		} else if probableValue < value {
			// find in the upper bound
			low = mid+1
		} else {
			high = mid-1
		}
	}

	return
}
