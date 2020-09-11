package selection

// Compare steps: (N-1)+(N-2)...
func SortInt(l []int) {
	size := len(l)
	for i := 0; i < size; i++ {
		min := i
		for j := i+1; j < size; j++ {
			if l[j] < l[min] {
				min = j
			}
		}
		// Swap, put the ministmum element
		// to its proper place
		l[i], l[min] = l[min], l[i]
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
