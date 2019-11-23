package main

import "fmt"

func HasDuplicate(elements []int) bool {
		tracker := make(map[int]struct{})
		steps := 0
		for i := 0; i < len(elements); i++ {
			steps++
			_, ok := tracker[elements[i]]
			if ok {
				return true
			}
			tracker[elements[i]] = struct{}{}
		}
		fmt.Println("steps:", steps)
		return false
}

func main() {
	input := []int{3, 2, 1}
	fmt.Println("Has duplicate?", HasDuplicate(input))
}
