package sort

func SortInt(elements []int, comp func(i, j int) bool) {
	size := len(elements)

	for i := 0; i < size-1; i++ { // size-1,since no need to select the last element
		for j := 0; j < size-i-1; j++ {
			if comp(elements[j], elements[j+1])	{
				elements[j], elements[j+1] = elements[j+1], elements[j]
			}
		}
	}
}

func SortIntReview(elements []int) {
	size := len(elements)

	for i := 0; i < size-1; i++ { // O(n)
		for j := 0; j < size-i-1; j++ { // O(n)
			if elements[j] > elements[j+1] { // Step 1:
				// swap
				elements[j], elements[j+1] = elements[j+1], elements[j]  // Step 2:
			}
		}
	}
}
