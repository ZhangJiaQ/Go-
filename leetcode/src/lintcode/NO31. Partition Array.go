package lintcode

import "fmt"

func partitionArray (nums []int, k int) int {
	// write your code here
	if len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1
	for left < right {

		fmt.Println(left, right)
		for left < right && nums[left] < k {
			left += 1
		}
		for left < right && nums[right] <= k {
			right -= 1
		}
		if left < right {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
		}
	}

	if nums[left] >= k {
		return left
	}
	return left + 1
}