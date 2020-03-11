package Array

func canJump(nums []int) bool {
	/*
	Given an array of non-negative integers, you are initially positioned at the first index of the array.
	Each element in the array represents your maximum jump length at that position.
	Determine if you are able to reach the last index.
	*/

	length := len(nums)

	if length < 2 {
		return true
	}

	// 入栈的时候计算 出栈的时候标为-1  遇到0  break

	tempStack := make([][]int, 0)

	index := 0
	for index < length - 1{
		// 标记为0的数字直接跳过
		if nums[index] == -1 {
			index += 1
			continue
		}

		if nums[index] == 0 {
			index += 1
			break
		}
		tempStack = append(tempStack, []int{0,nums[0]})
		// 开始跳跃，跳跃不到的均标0
		for len(tempStack) > 0 {
			// 入栈 应该改为深度优先遍历
			tempValue := tempStack[len(tempStack) - 1][1]
			tempIndex := tempStack[len(tempStack) - 1][0]
			startLength := len(tempStack)
			for tempValue > 0 {
				maxLength := tempIndex + tempValue
				if maxLength >= length - 1 {
					return true
				}	else if nums[maxLength] > 0 {
					tempStack = append(tempStack, []int {maxLength, nums[maxLength]})
					// 加break是深度 不加是广度，优先使用深度搜索
					break
				}
				tempValue -= 1
			}
			if len(tempStack) == startLength {
				// 该位置无论跳几轮都到不了终点，标记为-1
				nums[tempStack[len(tempStack) - 1][0]] = -1
				tempStack = tempStack[:len(tempStack) - 1]
			}
		}
	}
	return false
}