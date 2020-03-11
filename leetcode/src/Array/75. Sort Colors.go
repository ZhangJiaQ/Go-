package Array

import "fmt"

func sortColors(nums []int)  {

	/*
	大概就是  0 1 2排序 相同的放一起
	要求 in-place
	[0,2,2,1,1,0] - > [0,0,1,1,2,2]
	思路 两个指针 指向头部和尾部
	遍历数组  等于0则向头部指针交换数字  头部指针++
	          等于1则向尾部指针交换数字  尾部指针++
	*/
	zeroPoint := 0
	twoPoint := len(nums) - 1

	index := 0
	for index < twoPoint{
		if nums[index] == 1 {
			index += 1
		} else if nums[index] == 0 && index != zeroPoint {
			nums[index], nums[zeroPoint] = nums[zeroPoint], nums[index]
			zeroPoint += 1
			index += 1
		} else if nums[index] == 0 && index == zeroPoint{
			index += 1
			zeroPoint += 1
		} else if nums[index] == 2 {
			// 尾部指针交换
			for index < twoPoint && nums[twoPoint] == 2{
				twoPoint--
			}
			nums[index], nums[twoPoint] = nums[twoPoint], nums[index]
		}
	}
	if nums[len(nums)-1] == 0 {
		nums[len(nums)-1], nums[zeroPoint] = nums[zeroPoint], nums[len(nums)-1]
	}

	fmt.Println(nums)
}
