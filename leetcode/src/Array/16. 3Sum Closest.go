package Array

func threeSumClosest(nums []int, target int) int {
	/*
	Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target.
	Return the sum of the three integers. You may assume that each input would have exactly one solution.
	Example:
	Given array nums = [-1, 2, 1, -4], and target = 1.
	The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
	*/
	Quicksort(&nums)
	// 三根指针
	// [-5,-4,-3,-2,-1,0,1,2,3,4,5]  target 9
	oldSum := nums[0] + nums[1] + nums[len(nums)-1]
	oldDiff := abs(oldSum - target)
	for i, _ := range nums[0:len(nums) -1 ] {

		left, right := i + 1, len(nums) -1
		// 小于老diff时候  左指针右移动， 大于老diff得时候 右指针左移即可
		for left < right {
			newSum :=  nums[i] + nums[left] + nums[right]
			// 记录目前diff
			newDiff := abs(newSum - target)
			if newDiff < oldDiff {
				// 出现更小的diff了，记录
				oldDiff = newDiff
				oldSum = newSum
			}
			// 移动指针
			if newSum > target{
				right -= 1
			} else if newSum < target{
				left += 1
			} else {
				return target
			}
		}
	}

	return oldSum
}


func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}


func Quicksort(nums *[]int) {
	partition(nums, 0, len(*nums)-1)
}


func partition(nums *[]int, left, right int) {
	if left < right {
		mid := sort(nums, left, right)
		partition(nums, left, mid - 1)
		partition(nums, mid + 1, right)
	}
}


func sort(nums *[]int, left, right int) int {
	mid := (*nums)[left]

	for left < right {
		for left < right && (*nums)[right] >= mid {
			right -= 1
		}
		(*nums)[left] =  (*nums)[right]
		for left < right && (*nums)[left] <= mid {
			left += 1
		}
		(*nums)[right] =  (*nums)[left]
	}
	(*nums)[left] = mid
	return left
}