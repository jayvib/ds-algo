package insertion_sort

func Sort(data []int) {
	size := len(data)

	for i := 1; i < size; i++ {
		// Remove the value at
		// index i and store it
		// temporarily
		position := i
		tempValue := data[i]

		// Loop backward and check if
		// there's a greater value from
		// this position
		for position > 0 && data[position-1] > tempValue {
			data[position] = data[position-1]
			position--
		}

		// Put the temporary value to its correct position
		data[position] = tempValue
	}
}
