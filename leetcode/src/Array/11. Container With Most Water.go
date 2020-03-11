package Array

func maxArea(height []int) int {
	/*
	Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
	n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines,
	which together with x-axis forms a container, such that the container contains the most water.
	*/
	// 动态规划  假设ax 与 ay 的容积最长  那么 (y - x) * min(ax, ay)这个数最大

	// 使用 双指针来保证 y - x 最长
	// 在双指针移动的时候，移动两个指针中较小的值，来保证 min(ax, ay)最大
	// 记录每个数据，求出较大值保存，

	left := 0
	right := len(height) - 1
	maxAreaData := 0
	for left < right {
		tempArea := (right - left) * findMin(height[left], height[right])
		if tempArea > maxAreaData{
			maxAreaData  = tempArea
		}
		if height[left] > height[right] {
			right -= 1
		} else {
			left += 1
		}
	}


	return maxAreaData
}

func findMin(nums1, nums2 int) int {
	if nums1 < nums2 {
		return nums1
	} else {
		return nums2
	}
}