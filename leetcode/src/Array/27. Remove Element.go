package Array

func removeElement(nums []int, val int) int {
	/*
	Given an array nums and a value val, remove all instances of that value in-place and return the new length.
	Do not allocate extra space for another array,
	you must do this by modifying the input array in-place with O(1) extra memory.
	The order of elements can be changed. It doesn't matter what you leave beyond the new length.
	*/

	offset := 0
	index := 0

	for index < len(nums) {
		if nums[index] == val {
			offset += 1
		} else {
			nums[index-offset] = nums[index]
		}
		index += 1
	}

	offset = len(nums) - offset

	nums = nums[:offset]

	return offset
}