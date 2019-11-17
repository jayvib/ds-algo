package interpolation

// nums should be sorted.
func Search(nums []int, num int) bool {
	var mid, low, high int
	low, high = 0, len(nums)-1

	for nums[low] < num && nums[high] > num {
		// calculate mid using this formula
		mid = low + ((num-nums[low])*(high-low))/(nums[high]-nums[low])

		if nums[mid] < num {
			low = mid+1
		} else if nums[mid] > num {
			high = mid-1
		} else {
			return true
		}
	}

	if nums[low] == num {
		return true
	} else if nums[high] == num {
		return true
	}
	return false
}

func SearchReview(nums []int, num int) bool {
	var low, mid, high int
	low, high = 0, len(nums)-1

	for nums[low] < num && nums[high] > num {
		// calculate the mid
		mid = low + ((num-nums[low])*(high-low))/(nums[high]-nums[low])

		if nums[mid] < num {
			low = mid+1
		} else if nums[mid] > num { // find in lower partition
			high = mid-1
		} else {
			return true
		}
	}

	if nums[low] == num {
		return true
	} else if nums[high] == num {
		return true
	}
	return false
}

func SearchReview2(nums []int, num int) bool {
	var low, mid, high int
	low, high = 0, len(nums)-1

	for nums[low] < num && nums[high] > num {
		// calculate the interpolation
		mid = low + ((num-nums[low])*(high-low))/(nums[high]-nums[low])

		// find in higher partition
		if nums[mid] < num {
			low = mid+1

		// find in righ partition
		} else if nums[mid] > num {
				high = mid-1
		} else {
			return true
		}
	}

	if nums[low] == num {
		return true
	} else if nums[high] == num {
		return true
	}
	return true
}



