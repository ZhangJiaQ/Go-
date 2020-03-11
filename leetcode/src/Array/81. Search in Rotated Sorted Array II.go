package Array

func search(nums []int, target int) bool {
	/*
	81. Search in Rotated Sorted Array II
	Add to List

	Share
	Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

	(i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).

	You are given a target value to search. If found in the array return true, otherwise return false.
	*/
	// 1. 尾部与首部相等的就不管了 直接减去尾部长度
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		if target == nums[0]{
			return true
		} else {
			return false
		}
	}



	length := len(nums)
	for length >= 1 {
		if nums[length-1] == nums[0] {
			length -= 1
		} else {
			break
		}
	}

	for length >= 1 {
		if nums[length-1] == nums[length-2] {
			length -=1
		} else {
			break
		}
	}


	// 可以转换成首尾不相等的 但存在重复元素的数组

	left, right := 0, length - 1
	if right < left {
		// 就一个数的数组
		if target == nums[0]{
			return true
		} else {
			return false
		}
	}
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		} else if nums[right] > nums[mid] {
			// 右边有序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// 左边有序.
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false

}
