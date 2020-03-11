package Array

func searchInsert(nums []int, target int) int {

	/*
	Given a sorted array and a target value, return the index if the target is found.
	If not, return the index where it would be if it were inserted in order.
	You may assume no duplicates in the array.
	*/

	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2

		if mid > 0 && nums[mid] > target && nums[mid-1] < target {
			return mid
		}

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}

	}
	return right
}