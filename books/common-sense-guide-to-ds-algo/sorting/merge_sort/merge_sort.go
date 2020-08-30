package quicksort_sort

// ----Sort----
// Step 1: Check if the lower index is greater than or equal to upper index.
//				 If it is, then return
// Step 2: Calculate the middle index
// Step 3: Call merge sort recursively in the left bound array and right bound array
//  			 respectively.
// ----Merge----
// Step 4: Merge the arrays
//         Step 4.1: For two arrays, lower and upper. Compare each values
//									 from the two arrays and increment the temp array index.
//                   Cond 1: If the lower value is less than the upper value, then
//													 add it to the temporary array and increment the
//                           lower index
//                   Cond 2: If the lower value is greater than the upper value, then
//													 add it to the temporary array and increment the
//                           upper index
// Step 5: Exhaust the remaining lower array data, append it to the temp array
// Step 6: Exhaust the remaining upper array data, append it to the temp array

func merge(data, temp []int, lowerIndex, middleIndex, upperIndex int) {
	lowerStart := lowerIndex
	lowerStop := middleIndex
	upperStart := middleIndex+1
	upperStop := upperIndex
	count := lowerIndex

	for lowerStart <= lowerStop && upperStart <= upperStop {
		if data[lowerStart] < data[upperStart] {
			temp[count] = data[lowerStart]
			lowerStart++
		} else {
			temp[count] = data[upperStart]
			upperStart++
		}
		count++
	}

	// Step 5
	for lowerStart <= lowerStop {
		temp[count] = data[lowerStart]
		lowerStart++
		count++
	}

	// Step 6
	for upperStart <= upperStop {
		temp[count] = data[upperStart]
		upperStart++
		count++
	}

	for i := lowerIndex; i <= upperIndex; i++ {
		data[i] = temp[i]
	}
}

func mergeSort(data, temp []int, lowerIndex, upperIndex int) {
	if lowerIndex >= upperIndex {
		return
	}
	middleIndex := lowerIndex + (upperIndex - lowerIndex) / 2
	mergeSort(data, temp, lowerIndex, middleIndex)
	mergeSort(data, temp, middleIndex+1, upperIndex)
	merge(data, temp, lowerIndex, middleIndex, upperIndex)
}

func Sort(data []int) {
	mergeSort(data, make([]int, len(data)), 0, len(data)-1)
}
