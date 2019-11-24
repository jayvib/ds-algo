package insertion

func SortInt(l []int) {
	for i := 1; i < len(l); i++ {
		position := i
		tmp := l[i]
		for position > 0 && l[position-1] > tmp {
			l[position] = l[position-1]
			position--
		}
		l[position] = tmp
	}
}

func InsertionSortReview(elements []int) {
	for i := 1; i < len(elements); i++ { // from 2nd value to n
		position := i // the current position to remove
		tmp := elements[i] // store the current value of the empty position
		for position > 0 && elements[position-1] > tmp {
			elements[position] = elements[position-1]
			position--
		}
		elements[position] = tmp
	}
}

func InsertionSortReviewV2(elements []int) {
	for i := 1; i < len(elements); i++ {
		position := i
		tmp := elements[i]
		for position > 0 && elements[position-1] > tmp {
			// swap
			elements[position] = elements[position-1]
			position--
		}
		elements[position] = tmp
	}
}
