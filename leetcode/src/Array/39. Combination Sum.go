package Array

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	/*
	Given a set of candidate numbers (candidates) (without duplicates) and a target number (target),
	find all unique combinations in candidates where the candidate numbers sums to target.
	The same repeated number may be chosen from candidates unlimited number of times.
	Note:
	All numbers (including target) will be positive integers.
	The solution set must not contain duplicate combinations.
	*/
	Quicksort(&candidates)

	result := make([][]int, 0)
	temp := make([]int, 0)
	dfs(&candidates, &result, &temp, 0, target)

	return result
}

func dfs (candidates *[]int, result *[][]int, temp *[]int, index int, target int) {
	fmt.Println(temp)
	if target == 0 {
		copyResult := make([]int, 0)
		for _, v := range *temp {
			copyResult = append(copyResult, v)
		}
		*result = append(*result, copyResult)
	}
	for index < len(*candidates) && (*candidates)[index] <= target{
		*temp = append(*temp, (*candidates)[index])
		dfs(candidates, result, temp, index, target - (*candidates)[index])
		*temp = (*temp)[:len(*temp) - 1]
		index += 1
	}
}

