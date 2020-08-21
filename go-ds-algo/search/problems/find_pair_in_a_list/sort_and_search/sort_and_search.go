package brute_force

import "sort"

// Steps:
// 1. Sort all the elements in the list.
// 2. first index = 0, second index = size-1
// 3. compare the sum of two values if equal to
//    desired value.
// 4. if the sum is is greater than the value
//    then decrement second index
// 5. if the sum is less than the value then
//    increment the first index
func FindPair(ints []int, value int) bool {
	sort.Ints(ints)
	i, j := 0, len(ints)-1
	for i < j {
		sum := ints[i] + ints[j]
		switch {
		case sum > value:
			j--
		case sum < value:
			i++
		default:
			// Equal
			return true
		}
	}
	return false
}
