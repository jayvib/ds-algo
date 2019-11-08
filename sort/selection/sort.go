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