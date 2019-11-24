package selection

func SortInt(l []int) {
	size := len(l)
	for i := 0; i < size; i++ {
		lowestNumberIndex := i
		for j := i+1; j < size; j++ {
			if l[j] < l[lowestNumberIndex] {
				lowestNumberIndex = j
			}
		}

		if lowestNumberIndex != i {
			tmp := l[i]
			l[i] = l[lowestNumberIndex]
			l[lowestNumberIndex] = tmp
		}
	}
}

func SelectionSortInt(elements []int) {
	size := len(elements)
	for i := 0; i < size-1; i++ {
		max := 0
		for j := 1; j < size-i-1; j++ {
			if elements[j] > max {
				max = j
			}
		}
		// put the max value index to the last element
		elements[size-i-1], elements[max] = elements[max], elements[size-i-1]

	}
}
