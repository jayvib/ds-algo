package lumoto_partition

func SeparateOddEven(data []int) {
	size := len(data)
	j := -1
	for i := 0; i < size; i++ {
		// If even
		if data[i]%2 == 0 {
			j++
			// Swap
			swap(data, i, j)
		}
	}
}

func swap(data []int, i int, j int) {
	data[i], data[j] = data[j], data[i]
}
