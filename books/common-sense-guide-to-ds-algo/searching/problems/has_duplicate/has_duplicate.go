package has_duplicate

import "fmt"

func HasDuplicate(data []int) (found bool) {
	return hasDuplicateMap(data)
}

// Time Complexity: O(N)
func hasDuplicateMap(data []int) bool {
	steps := 0
	existingNumberCounter := make(map[int]int)
	for i := 0; i < len(data); i++ {
		steps++
		if _, ok := existingNumberCounter[data[i]]; ok {
			return true
		}
		existingNumberCounter[data[i]]++
	}
	fmt.Println(steps)
	return false
}

// Time Complexity: O(N^2)
func hasDuplicateNested(data []int) bool {
	steps := 0
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			steps++
			if data[i] == data[j] {
				return true
			}
		}
	}
	fmt.Println(steps)
	return false
}

