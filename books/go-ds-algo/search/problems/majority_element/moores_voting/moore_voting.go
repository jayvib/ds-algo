package moores_voting

import "fmt"

func FindMajority(data []int) (int, bool) {
	majIndex := 0
	count := 1
	size := len(data)
	for i := 1; i < size; i++ {
		if data[majIndex] == data[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			majIndex = i
			count = 1
		}
	}
	candidate := data[majIndex]
	count = 0
	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		return data[majIndex], true
	}
	fmt.Println("MajorityDoesNotExist")
	return 0, false
}
