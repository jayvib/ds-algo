package bubble_sort

func Sort(data []int) {
	size := len(data)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func SortImproved(data []int) {
	isSwapped := true
	size := len(data)
	for i := 0; i < size && isSwapped; i++ {
		isSwapped = false
		for j := 0; j < size-i-1 ; j++ {
			if data[j+1] < data[j] {
				isSwapped = true
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
