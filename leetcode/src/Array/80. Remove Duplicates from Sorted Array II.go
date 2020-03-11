package Array

func removeDuplicates(nums []int) int {

		/*
		Given a sorted array nums, remove the duplicates in-place such that duplicates appeared at most twice and return the new length.
		Do not allocate extra space for another array,
		you must do this by modifying the input array in-place with O(1) extra memory.
		*/
		if len(nums) == 0 {
			return 0
		}

		count := 1
		nowNum := nums[0]
		nowCount := 1
		offset := 0
		index := 1
		for index < len(nums) {
			// 00 1111 2 33
			if nowNum == nums[index] {
				if nowCount == 1 {
					count += 1
					nowCount += 1
				} else {
					offset += 1
				}
			} else {
				nowNum = nums[index]
				nowCount = 1
				count += 1
			}
			nums[index-offset] = nums[index]
			index += 1
		}

		nums = nums[:len(nums)-offset]

		return count
}