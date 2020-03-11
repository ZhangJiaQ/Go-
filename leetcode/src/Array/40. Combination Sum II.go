package Array
func combinationSum2(candidates []int, target int) [][]int {


	// 与之前不同，这次每个元素只能用一次

	// dfs得基础上，每次循环Index+1 不走自己的老数字
	Quicksort(&candidates)
	result := make([][]int, 0)
	temp := make([]int, 0)

	dfsWithoutSelf(&candidates, &result, &temp, 0, target)

	return result
}


func dfsWithoutSelf (candidates *[]int, result *[][]int, temp *[]int, index, target int) {

	if target == 0 {
		copyTemp := make([]int, 0)
		for _, v := range *temp{
			copyTemp = append(copyTemp, v)
		}
		*result = append(*result, copyTemp)
	}

	// 使用tempIndex的原因就是，在同一列进行遍历的时候遇到相同的则跳过，但是深度不同的遍历不受影响。，
	tempIndex := index

	for tempIndex < len(*candidates) && (*candidates)[tempIndex] <= target{
		if tempIndex > index && (*candidates)[tempIndex-1] == (*candidates)[tempIndex]{
			tempIndex += 1
			continue
		}
		*temp = append(*temp, (*candidates)[tempIndex])
		dfsWithoutSelf(candidates, result, temp, tempIndex + 1, target-(*candidates)[tempIndex])
		*temp = (*temp)[:len(*temp)-1]
		tempIndex += 1
	}
}