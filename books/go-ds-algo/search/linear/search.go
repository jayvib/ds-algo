package linear

// Time Complexity: O(n) -> Worst Case
//                  O(1) -> Best case
func SearchUnsorted(value int, data []int) bool {
	for _, d := range data {
		if value == d {
			return true
		}
	}
	return false
}

// Time Complexity:
// O(n)   -> Worst Case
// O(n/2) -> Average Worst Case
// O(1)   -> Best Case
func SearchSorted(value int, data []int) bool {
	for i := 0; i < len(data); i++ {
		if value == data[i] {
			return true
		}

		if value < data[i] {
			return false
		}
	}
	return false
}
