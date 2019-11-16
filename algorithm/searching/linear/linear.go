package linear

// Time Complexity: O(n)
func Search(items []int, find int) bool {
	for i := range items {
		if items[i] == find {
			return true
		}
	}
	return false
}

