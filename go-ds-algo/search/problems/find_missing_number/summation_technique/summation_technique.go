package summation_technique

func FindMissingElement(ints []int) (value int, found bool) {
	size := len(ints)
	total := (size+1)*(size+2)/2
	for i := 0; i < size; i++ {
		total -= ints[i]
	}

	if total == size+1 {
		return 0, false
	}

	value = total
	found = true
	return
}
