package Array

func remove1Duplicates(nums []int) int {

	/*
	Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
	Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
	*/

	// 意思就是 去除数组中的重复元素，要求在O 1 的空间复杂度
	// 用两个指针 一个记录重复的偏移量 一个用于移动对象

	offset := 0
	index := 1
	for index < len(nums) {
		if nums[index] == nums[index - 1] {
			offset += 1
		} else {
			// 不相等
			nums[index-offset] = nums[index]
		}
		//  0 0 1 1 1 2 3
		index += 1
	}
	offset = len(nums) - offset
	nums = nums[:offset]
	return  offset
}
