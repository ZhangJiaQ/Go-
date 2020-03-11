package Array

func largestRectangleArea(heights []int) int {
	/*
	Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.
	*/

	// 寻找局部峰值 然后向前遍历


	if len(heights) == 0 {
		return 0
	}

	heights = append(heights, 0)
	MaxArea := 0
	index := 0
	for index < len(heights) - 1 {
		if heights[index] >  heights[index+1] {
			// 发现局部最大值
			tempIndex := index
			tempMinNum := heights[index]
			for tempIndex >= 0 {
				if tempMinNum > heights[tempIndex]{
					tempMinNum = heights[tempIndex]
				}
				if tempMinNum*(index-tempIndex+1) > MaxArea {
					MaxArea = tempMinNum*(index-tempIndex+1)
				}
				tempIndex--
			}
		}
		index += 1
	}
	heights = heights[:len(heights)-1]

	return MaxArea



}