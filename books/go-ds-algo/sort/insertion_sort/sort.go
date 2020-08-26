package insertion_sort

func Sort(data []int) {
	// 3 1 2 5
	for i := 1; i < len(data); i++ {
		selection := data[i]
		var j int
		for j = i; j > 0 &&  selection < data[j-1]; j-- {
			data[j], data[j-1] = data[j-1], data[j]
		}
		data[j] = selection
	}
}
