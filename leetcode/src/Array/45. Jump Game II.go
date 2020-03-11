package Array

func jump(nums []int) int {
	// 贪心算法，不断更新目标

	// 找到第一个可以到goal位置的位置X 然后找第一个可以到达位置X的参数X  直到X 等于 0

	goalIndex := len(nums) - 1
	jump := 0
	for goalIndex != 0 {
		for index, value := range nums {
			if index + value >= goalIndex {
				goalIndex = index
				jump ++
				break
			}
		}
	}


	return jump
}