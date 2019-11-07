package bubble

func SortInt(l []int) {
	unsortedUntilIndex := len(l)-1
	isSorted := false // assume first the the list is not sorted
	for !isSorted { // while the list is not sorted when keep the loop
		isSorted = true // assume that this passthrough is sorted
		for i := 0; i < unsortedUntilIndex; i++ { // compare the adjacent values. and move the greatest value to its designated position
			if l[i] > l[i+1] { // compare the elements
				l[i], l[i+1] = l[i+1], l[i] // swap
				isSorted = false // set variable not sorted.
			}
		}
		unsortedUntilIndex--
	}
}
