package Array

func subsets(nums []int) [][]int {

	/*
	Given a set of distinct integers, nums, return all possible subsets (the power set)
	Note: The solution set must not contain duplicate subsets.
	*/
	result := make([][]int, 0)
	length := len(nums)
	temp := make([]int, 0)

	index := 0
	for index <= length {
		bfsSubset(&result, &nums, &temp, index, 0)
		index += 1
	}

	return result
}

func bfsSubset(result *[][]int, nums *[]int, temp *[]int,  length int, start int) {
	if length == 0 {
		copyTemp := make([]int, len(*temp))
		copy(copyTemp, *temp)
		*result = append(*result, copyTemp)
	}

	index := start
	for index < len(*nums) {
		*temp = append(*temp, (*nums)[index])
		bfsSubset(result, nums, temp, length-1, index+1)
		*temp = (*temp)[:len(*temp)-1]
		index += 1
	}

}

