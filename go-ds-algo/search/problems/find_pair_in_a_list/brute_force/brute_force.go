package brute_force

// FindPair finds the two elements that is equal to
// the value.
// Time complexity is O(n^2)
// Space complexity is O(1)
func FindPair(ints []int, value int) bool {
	// First loop select the item to compare
	size := len(ints)
	for i := 0; i < size; i++ {
		// Second loop will do comparing
		for j := i+1; j < size; j++ {
			if sum(ints[i], ints[j]) == value {
				return true
			}
		}
	}
	return false
}

func sum(x int, y int) int {
	return x + y
}
