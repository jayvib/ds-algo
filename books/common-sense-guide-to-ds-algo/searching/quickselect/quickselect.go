package quickselect

import "github.com/jayvib/ds-algo/books/common-sense-guide-to-ds-algo/partition"

func QuickSelect(data []int, kthLowestIndexValue int) int {
	return quickSelect(data, kthLowestIndexValue, 0, len(data)-1)
}

// Step 1: Do partion of the array and get the pivot index
// Step 2: Compare the kth lowest index with the pivot index:
//         Cond 1: If the kth lowest index is lesser then the pivot index,
//								 then do a recursive call with a new rightIndex of pivotIndex-1
//         Cond 2: If the kth lowest index is greater than the pivot index,
// 							   then do a recursive call with a new leftIndex of pivotIndex+1
//         Cond 3: If equal then return the value of the pivot index.
// Step 3: When the when the subarray has one cell return the value of that cell
func quickSelect(data []int, kthLowestIndexValue, leftIndex, rightIndex int) int {

	if (rightIndex-leftIndex) <= 0 {
		return data[leftIndex]
	}

	pivotIndex := partition.Partition(data, leftIndex, rightIndex)

	if kthLowestIndexValue < pivotIndex {
		return quickSelect(data, kthLowestIndexValue, leftIndex, pivotIndex-1)
	} else if kthLowestIndexValue > pivotIndex {
		return quickSelect(data, kthLowestIndexValue, pivotIndex+1, rightIndex)
	} else {
		return data[pivotIndex]
	}
}
