package quicksort

import "math/rand"

func Sort(l []int) []int {
	if len(l) < 2 { // this will be the base case. Then the subarray has zero or one element
		return l
	}

	// define the left and right pointers
	left, right := 0, len(l)-1

	// define the pivot
	pivot := rand.Int() % len(l)

	// move the pivot at the last element

	l[pivot], l[right] = l[right], l[pivot] // so the pivot is already at the last index

	// move the numbers that are less than the pivot
	// to the left side.
	for i := range l {
		if l[i] < l[right] {
			l[left], l[i] = l[i], l[left]
			left++
		}
	}

	// swap the current pivot to the current right pointer
	l[left], l[right] = l[right], l[left]

	// Sort the values that are less than the current pivot
	Sort(l[:left])

	// Sort the value that are greater than the current pivot
	Sort(l[left+1:])
	return l
}
