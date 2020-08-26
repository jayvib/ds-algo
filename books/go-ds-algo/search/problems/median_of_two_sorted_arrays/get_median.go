package median_of_two_sorted_arrays

import "fmt"

// Use merge procedure of merge sort.
// Keep track of count while comparing elements of two arrays.
// If count becomes n(For 2n elements), we have reached the median.
// Take the average of the elements at indexes n-1 and n in the merged array.
// See the below implementation.

func GetMedian(arr1, arr2 []int, n int) int {
	var i, j int

	m1, m2 := -1, -1

	// arr1 := []int{1, 12, 15, 26, 38}
	// arr2 := []int{2, 13, 17, 30, 45}
	for count := 0; count <= n; count++ {
		if i == n {
			m1 = m2
			m2 = arr2[0]
			break
		} else if j == n {
			m1 = m2
			m2 = arr1[0]
			break
		}

		if arr1[i] <= arr2[j] {
			// Get the lowest value from arr1
			// which is lowest value from arr2
			m1 = m2 // previous value of m2, which is value from arr2
			m2 = arr1[i]
			fmt.Printf("1 %01d, %01d\n", m1, m2)
			i++
		} else {
			m1 = m2 // previous value of m2, which is value from arr1
			m2 = arr2[j]
			fmt.Printf("2 %01d, %01d\n", m1, m2)
			j++
		}
	}

	return (m1 + m2) / 2
}
