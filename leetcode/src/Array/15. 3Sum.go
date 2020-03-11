package Array

import "fmt"

func threeSum(nums []int) [][]int {
	// 思路就是 先排序

	result := make([][]int, 0)
	// QuickSort(&nums)
	fmt.Println(nums)
	for index, value := range nums{
		// 然后三根指针， 第一个数P1为外层循环，需要为负
		if value > 0 {
			return result
		}
		if index > 0 && nums[index] == nums[index - 1]{
			continue
		}

		// 第二个数第三个数两根指针左右互博
		needSum := 0 - value
		left, right := index + 1, len(nums) -1

		for left < right {
			if nums[left] + nums[right] == needSum{
				temp := []int{value, nums[left], nums[right]}
				// 过滤掉重复的结果
				result = append(result, temp)
				right -= 1
				left += 1
				for left < right && nums[right+1] == nums[right]{
					right -= 1
				}
				for left < right && nums[left-1] == nums[left]{
					left += 1
				}
			} else if nums[left] + nums[right] > needSum {
				right -= 1
				for left < right && nums[right+1] == nums[right]{
					right -= 1
				}
			} else {
				left += 1
				for left < right && nums[left-1] == nums[left]{
					left += 1
				}

			}
		}

	}

	return result
}


//func quickSort(nums *[]int) {
//	partition(nums, 0, len(*nums)-1)
//}
//
//
//func partition(nums *[]int, left, right int) {
//	if left < right{
//		fmt.Println(left, right)
//		mid := sort(nums, left, right)
//		partition(nums, left, mid - 1)
//		partition(nums, mid + 1, right)
//	}
//}
//
//func sort(nums *[]int, left, right int) int {
//	mid := (*nums)[left]
//	for left < right {
//		for left < right && (*nums)[right] >= mid {
//			// 当右边数比 mid基准数 大的时候，右指针左移
//			right -= 1
//		}
//		// 目前右指针的数 已经比 mid基准数小,将右指针的数字移到 mid基准数 的位置
//		(*nums)[left] = (*nums)[right]
//		for left < right && (*nums)[left] <= mid {
//			// 当右边数比 mid基准数 小的时候，左指针右移
//			left += 1
//		}
//		// 目前左指针的数 已经比 mid基准数大,将左指针的数字移到 mid基准数 的位置
//		(*nums)[right] = (*nums)[left]
//	}
//	// 跳出循环时 左指针 指的即为 mid基准数应该在的位置
//	(*nums)[left] = mid
//	return left
//}