package bubble_sort

func Sort(data []int) {
	size := len(data)
	unsortedUntilIndex := size-1
	isSorted := false

	for !isSorted {
		isSorted = true

		// i -> N
		for i := 0; i < unsortedUntilIndex; i++ {
			if data[i] > data[i+1] { // O(1)
				// Swap
				data[i], data[i+1] = data[i+1], data[i] // STEP: Swap with O(1)
				isSorted = false // O(1)
			}
		}
		unsortedUntilIndex-- // O(1)
	}
}
