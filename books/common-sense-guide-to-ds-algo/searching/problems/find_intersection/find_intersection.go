package find_intersection

func FindIntersection(arr1, arr2 []int) (intersection []int) {
	// Step 1: Compare
	// Step 2: Insert
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				intersection = append(intersection, arr1[i])
				break
			}
		}
	}
	return
}
