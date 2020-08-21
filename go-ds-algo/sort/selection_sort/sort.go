package selection_sort

import "fmt"

func Sort(data []int) {
	size := len(data)
	var min int
	for i := 0; i < size; i++ {
		min = i
		for j := i+1; j < size; j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		fmt.Println(data)
		data[i], data[min] = data[min], data[i]
	}
}
