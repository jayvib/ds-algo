package main

import "sort"

// Problem:
// Given a list of n number, find two elements such that their sum is equal to "value"

func FindPairBrute(in []int, value int) [][2]int {
	// Steps:
	// 1. Select a value from the slice.
	// 2. Iterate throughout the list and add the current iterated value
	//	  to the selected value and compare if equal to value.
	// 3. When equal, add the result to the output list.

	output := make([][2]int, 0)
	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			if (in[i] + in[j]) == value {
				output = append(output, [2]int{in[i], in[j]})
			}
		}
	}

	return output
}

func FindPairSorted(in []int, value int) [][2]int {
	size := len(in)
	first, second := 0, size-1
	sort.Ints(in)
	out := make([][2]int, 0)
	for first < second {
		curr := in[first] + in[second]
		if curr == value {
			out = append(out, [2]int{in[first], in[second]})
		} else if curr < value {
			first++
		} else if curr > value {
			second--
		}
	}
	return out
}

func main() {

}
