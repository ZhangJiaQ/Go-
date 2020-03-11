package Array

func searchRange(nums []int, target int) []int {
	/*
	Given an array of integers nums sorted in ascending order,
	find the starting and ending position of a given target value.
	Your algorithm's runtime complexity must be in the order of O(log n).
	If the target is not found in the array, return [-1, -1].
	*/
	result := make([]int, 2)
	result[0], result[1] = -1, -1
	left, right := 0, len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			result[0], result[1] = mid, mid
			for result[0] > 0 && nums[result[0] - 1] == target {
				result[0] -= 1
			}
			for result[1] < len(nums) - 1 &&  nums[result[1] + 1] == target {
				result[1] += 1
			}
			return result
		} else if nums[mid] < target{
			left = mid + 1
		} else {
			right = mid  - 1
		}
	}


	return result
}