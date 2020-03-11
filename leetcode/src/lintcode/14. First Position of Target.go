package lintcode

/**
 * @param nums: The integer array.
 * @param target: Target to find.
 * @return: The first position of target. Position starts from 0.
 */
func binarySearch (nums []int, target int) int {
	// write your code here
	// 二分搜索
	result := -1

	left := 0
	right := len(nums) - 1

	for left < right - 1 {
		mid := (left + right) / 2
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid
		}
	}

	if nums[left] == target{
		return left
	}

	if nums[right] == target{
		return right
	}

	return result
}
