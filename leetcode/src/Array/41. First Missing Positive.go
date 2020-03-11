package Array

import "fmt"

func firstMissingPositive(nums []int) int {

	// 找到最小的 缺失的正整数

	/*
	Input: [1,2,0]
	Output: 3

	Input: [3,4,-1,1]
	Output: 2

	Input: [7,8,9,11,12]
	Output: 1

	*/

	length := len(nums)
	// 思路，将
	for index, _ := range nums {
		if nums[index] == index + 1 {
			continue
		}
		for nums[index] < length && nums[index] > 0  && nums[index] != nums[nums[index] - 1] {
			swap(&nums[index], &nums[nums[index]-1])
		}
	}
	fmt.Println(nums)
	index := 0
	for index < length {
		if nums[index] != index + 1 {
			break
		}
		index += 1
	}

	return index + 1

}

func swap(left, right *int ) {
	temp := *left
	*left = *right
	*right = temp
}


func firstMissingPositiveBad(nums []int) int {

	// 找到最小的 缺失的正整数

	/*
	Input: [1,2,0]
	Output: 3

	Input: [3,4,-1,1]
	Output: 2

	Input: [7,8,9,11,12]
	Output: 1

	*/

	left, right := 1, nums[0]
	min := nums[0]
	max := nums[0]
	if nums[0] < 0 {
		max = 0
	}

	for _, v := range nums {
		if v > 0 {
			if v > max {
				max = v
			}
			if v <= left {
				left = v + 1

			} else if v <= min {
				min = v
				right = min - 1
			}

		}
		if max > 0 && left > right {
			left = max + 1
			right = max + 1
		}
	}
	result := Min(left, right)
	return result

}

func Min(left, right int) int {
	if left < right{
		return left
	} else {
		return right
	}
}