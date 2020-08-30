package partition

func Partition(data []int, leftCursor, rightCursor int) (partitionIndex int) {
	// Pre-step 1: Choose a pivot, use the right most element as pivot
	//
	// For each pass-through
	// Step 1: Move the left cursor where the value is greater than the pivot.
	// Step 2: Move the right cursor where the value is less than the pivot.
	// Step 3: If left cursor index is equal or greater than the right then:
	//         Cond 1(false): break the pass-through
	//         Cond 2: swap the left and right cursor value
	//
	// Post-step 1: Swap the pivot with the left cursor index.

	// Pre-step 1: Choose a pivot, use the right most element as pivot
	pivotCursor := rightCursor
	pivot := data[pivotCursor]

	rightCursor--

	for {

		// Step 1: Move the left cursor where the value is greater
		// than the pivot.
		for data[leftCursor] < pivot {
			leftCursor++
		}

		// Step 2: Move the right cursor where the value is less than
		// the pivot.
		for data[rightCursor] > pivot {
			rightCursor--
		}

		// Step 3: Cond 1
		if leftCursor >= rightCursor {
			break
		}

		// Step 3: Cond 2
		data[leftCursor], data[rightCursor] = data[rightCursor], data[leftCursor]
	}

	// Post-step 1
	data[pivotCursor], data[leftCursor] = data[leftCursor], data[pivotCursor]
	return leftCursor
}
