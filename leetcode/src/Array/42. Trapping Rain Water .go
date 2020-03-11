package Array

func trap(height []int) int {
	/*
	Given n non-negative integers representing an elevation map where the width of each bar is 1,
	compute how much water it is able to trap after raining.
	*/

	// 动态规划，我们遍历heght数组
	// 遇到 height[index] 小于height[index-1] 与 height[index+1]的时候开始计算储水量

	if len(height) < 3{
		return 0
	}
	length := len(height)
	left, right := 0, length - 1
	lBound, rBound := 0, 0

	result := 0

	for left < right {
		lBound = max(height[left], lBound)
		rBound = max(height[right], rBound)

		if lBound <= rBound{
			result += lBound - height[left]
			left += 1
		} else {
			result += rBound - height[right]
			right -= 1
		}
	}

	return result
}


func max(left, right int) int {
	if left > right {
		return left
	} else {
		return right
	}
}