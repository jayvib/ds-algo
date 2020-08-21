package jump_search

import "math"

func Search(data []int, number int) (index int, found bool) {
	size := len(data)
	step := int(math.Sqrt(float64(size)))

	prev := 0
	// find the first value that is greater than the search value
	for data[int(math.Min(float64(step), float64(size)))-1] < number {
		prev = step
		step += int(math.Sqrt(float64(size)))
		if prev >= size {
			return -1, false
		}
	}

	// Do linear search or binary search
	for data[prev] < number{
		prev++
		if prev == int(math.Min(float64(step), float64(size))) {
			return -1, false
		}
	}

	if data[prev] == number {
		return prev, true
	}

	return -1, false
}

func search(data []int, x int) (index int, found bool) {
	size := len(data)

	// Calculate equal steps
	step := int(math.Sqrt(float64(size)))

	// Find the first index which value is greater than
	// the search value
	prev := 0
	for data[int(math.Min(float64(step), float64(size)))-1] < x {
		prev = step
		step += int(math.Sqrt(float64(size)))
		if prev >= size {
			return -1, false
		}
	}
	// Do linear search
	for data[prev] != x {
		prev++
		if prev >= int(math.Min(float64(step), float64(size))) {
			return -1, false
		}
	}

	if data[prev] == x {
		return prev, true
	}
	return -1, false
}