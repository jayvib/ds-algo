package exponential

import "math"

func binarySearch(data []int, low, high, x int) (index int) {
	var mid int
	for low < high {
		mid = low + (high-low)/2
		if data[mid] == x {
			return mid
		} else if data[mid] < x {
			low = mid+1
		} else {
			high = mid-1
		}
	}
	return -1
}

func Search(arr []int, x int) (index int) {
	return exponentialSearch(arr, x)
}

// Exponential Search
func exponentialSearch(arr []int, x int) int {
	if arr[0] == x {
		return 0
	}

	i := 1
	for i < len(arr)-1 && arr[i] <= x {
		i *= 2
	}

	return binarySearch(arr, i/2, int(math.Min(float64(i), float64(len(arr)-1))), x)
}

