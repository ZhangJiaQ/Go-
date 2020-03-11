package lintcode

/**
 * @param nums: An integer array
 * @return: nothing
 */
func recoverRotatedSortedArray (nums *[]int)  {
	// write your code here
	index := 0
	// 找到应该旋转的位置
	for {
		if index == len(*nums) {
			return
		}
		if (*nums)[index] > (*nums)[index+1] {
			reverse(nums, 0, index)
			reverse(nums, index+1, len(*nums)-1)
			reverse(nums, 0, len(*nums)-1)
			break
		}
		index += 1
	}

}

func reverse(nums *[]int, left int, right int) {
	for left < right {
		swapNum(nums, left, right)
		left += 1
		right -= 1
	}
}


func swapNum(nums *[]int, left int, right int) {
	temp := (*nums)[left]
	(*nums)[left] = (*nums)[right]
	(*nums)[right] = temp
}
