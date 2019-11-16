package quicksort

import (
	"math/rand"
)

func QuickSort(l []int) []int {
	if len(l) < 2 { // this will be the base case. Then the subarray has zero or one element
		return l
	}

	// define the left and right pointers
	left, right := 0, len(l)-1

	// define the pivot select a random pivot
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

	// QuickSort the values that are less than the current pivot
	QuickSort(l[:left])

	// QuickSort the value that are greater than the current pivot
	QuickSort(l[left+1:])
	return l
}

func quickSort(l []int) []int {

	// base case
	size := len(l)
	if size < 2 {
		return l
	}

	// determine the left and right pointers
	left, right := 0, size-1

	// determine the pivot index
	pivot := rand.Int() % size

	// move the pivot to the last element	of the array.
	l[right], l[pivot] = l[pivot], l[right]

	// move the values that is less than the pivot to the left side
	for i := 0; i < size; i++ {
		if l[i] < l[right] {
			// put it to the left partition
			l[left], l[i] = l[i], l[left]
			left++
		}
	}

	// swap the current pivot to the upper boundary to the left side partition
	l[left], l[right] = l[right], l[left]

	quickSort(l[:left])
	quickSort(l[left+1:])
	return l
}


func QuickSortReview(array []int) []int {
	size := len(array)
	if size < 2 {
		return array
	}

	// determine the left and right index
	left, right := 0, size-1

	// pick a random pivot index
	pivot := rand.Int() % size

	// put the pivot value to the last array
	array[right], array[pivot] = array[pivot], array[right]

	// group the value that are less than the pivot
	// to the left side of the array
	for i := 0; i < size; i++ {
		if array[i] < array[right] {
			array[left], array[i] = array[i], array[left]
			left++
		}
	}

	// make a partition. Put the pivot
	// to the correct position
	array[right], array[left] = array[left], array[right]

	// Quick Sort the left group array
	QuickSortReview(array[:left])

	// Quick Sort the right group array
	QuickSortReview(array[left+1:])
	return array
}