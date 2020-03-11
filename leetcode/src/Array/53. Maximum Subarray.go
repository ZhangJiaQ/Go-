package Array

func maxSubArray(nums []int) int {
	/*
	Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
	*/

	tempArray := make([]int, len(nums))
	tempArray[0] = nums[0]
	temp := nums[0]

	// -2,1,-3,4,-1,2,1,-5,4
	index := 1
	for index < len(nums) {
		tempArray[index] = Max(nums[index] + tempArray[index-1], nums[index])
		index += 1
	}
	temp = tempArray[0]
	index = 1
	for index < len(tempArray) {
		if tempArray[index] > temp{
			temp = tempArray[index]
		}
		index += 1
	}
	return temp

}

func Max(left, right int) int {
	if left > right {
		return left
	} else {
		return right
	}
}
