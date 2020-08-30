package quicksort_sort

import "github.com/jayvib/ds-algo/books/common-sense-guide-to-ds-algo/partition"

func Sort(data []int) {
	quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, leftIndex, rightIndex int) {
	if (rightIndex - leftIndex) <= 0 {
		return
	}

	// Partition
	pivotIndex := partition.Partition(data, leftIndex, rightIndex)

	// Left boundary
	quickSort(data, leftIndex, pivotIndex-1)

	// Right boundary
	quickSort(data, pivotIndex+1, rightIndex)
}
