package brute_force

func OccurenceCount(arr []int, x int) int {
	count := 0
	for i := range arr {
		if arr[i] == x {
			count++
		}
	}
	return count
}
