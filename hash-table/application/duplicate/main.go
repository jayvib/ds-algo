package main

import "fmt"

func hasDuplicate(l []int) bool {
	existingValues := make(map[int]interface{}) // To track the values in the list
	for i := 0; i < len(l); i++ {
		if _, ok := existingValues[l[i]]; !ok {
			existingValues[l[i]] = struct{}{}
			continue
		}
		return true
	}
	return false
}

func main() {
	input := []int{1, 2, 3}
	fmt.Println(hasDuplicate(input))
}
