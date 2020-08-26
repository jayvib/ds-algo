package brute_force

func FindMaxRecurring(ints []int) int {
	size := len(ints)
	max := ints[0]
	maxCount := 1
	count := 1
	// Two for loops. The first loop
	// selects an item and the second loop
	// will find the same item.
	for i := 0; i < size; i++ {
		for j := i+1; j < size; j++ {
			if ints[i] == ints[j] {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			max = ints[i]
		}
	}
	return max
}
