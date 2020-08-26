package brute_force

//	Algorithm: segregateEvenOdd()
//	1) Initialize two index variables left and right:
//	left = 0,  right = size -1
//	2) Keep incrementing left index until we see an odd number.
//	3) Keep decrementing right index until we see an even number.
//	4) If lef < right then swap arr[left] and arr[right]

// Input  = {12, 34, 45, 9, 8, 90, 3}
// Output = {12, 34, 90, 8, 9, 45, 3}
func SeparateOddEven(data []int) {
	size := len(data)
	left := 0
	right := size-1

	for left < right {
		// Keep incrementing left index until we see an odd number
		for (data[left]%2 == 0) && (left < right){
			left++
		}
		// Keep decrementing right index until we see an even number
		for (data[right]%2 != 0) && (left < right) {
			right--
		}
		if left < right {
			// Swap
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}
	}
}
