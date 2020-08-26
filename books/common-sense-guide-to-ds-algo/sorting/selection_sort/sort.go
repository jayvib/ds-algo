package selection_sort

// Efficiency:
//
func Sort(data []int) {
	// Step 1: Select the item
	// Step 2: Compre the selected item to the right
	//         to find the index of the lowest value
	//         so far.
	// Step 3: Swap the lowest index with the selected item
	for i := 0; i < len(data); i++ { // Each passthrough
		lowestIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[lowestIndex] {
				lowestIndex = j
			}
		}

		// swap
		// No need to swap itself.
		if lowestIndex != i {
			data[i], data[lowestIndex] = data[lowestIndex], data[i]
		}
	}

}
