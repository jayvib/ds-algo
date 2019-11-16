package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createArray(n int) (out []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		out = append(out, rand.Intn(99999)-rand.Intn(99999))
	}
	return
}

func MergeSort(array []int) []int {
	if len(array) < 2 { // base case is when the length of an array is lesser than 2
		return array
	}
	var middle int
	middle = len(array) / 2 // calculate the middle index
	return JoinArray(MergeSort(array[:middle]), MergeSort(array[middle+1:]))
}

func JoinArray(leftArr []int, rightArr []int) []int {
	var i, j int
	n := len(leftArr) + len(rightArr) // the total number of items
	arr := make([]int, n) // create a container for the sorted items

	for k := 0; k < n; k++ {
		if i > len(leftArr)-1 && j <= len(rightArr)-1 { // Process the remaining item from right array
			arr[k] = rightArr[j]
			j++
		} else if j > len(rightArr)-1 && i <= len(leftArr)-1 { // Process the remaining item from left array
			arr[k] = leftArr[i]
			i++
		} else if leftArr[i] < rightArr[j] { // if left value is lesser than the right value then add the i value to arr
			arr[k] = leftArr[i]
			i++ // move cursor to next item.
		} else {
			arr[k] = rightArr[j] // if right value is lesser than the left value
			j++ // move cursor to next item.
		}
	}
	return arr
}

func MergeSortReview(nums []int) []int {
	// define the base case
	if len(nums) < 2 {
		return nums
	}

	// calculate the middle index
	var middle int
	middle = len(nums) / 2

	// partition the array and use recursion
	return JoinArrayReview(MergeSort(nums[:middle]), MergeSort(nums[middle+1:]))
}

func JoinArrayReview(leftArr, rightArr []int) []int {
	total := len(leftArr)+len(rightArr)
	merged := make([]int, total)
	var i, j, k int

	for k = 0; k < total; k++ {
		// process when the left side is exhausted
		if i > len(leftArr)-1 && j <= len(rightArr)-1 {
			merged[k] = rightArr[j]
			j++

		// process when the right side is exhausted
		} else if j > len(rightArr)-1 && i <= len(leftArr)-1 {
			merged[k] = leftArr[i]
			i++

		// when the left array value is lesser than the right array value
		} else if leftArr[i] < rightArr[j] {
			merged[k] = leftArr[i]
			i++

		// when the right array value is lesser than the left array value
		} else if leftArr[i] > rightArr[j] {
			merged[k] = rightArr[j]
			j++
		}
	}
	return merged
}

func main() {
	arr := createArray(100)
	sortedArr := MergeSortReview(arr)
	for _, v := range sortedArr {
		fmt.Println(v)
	}
}