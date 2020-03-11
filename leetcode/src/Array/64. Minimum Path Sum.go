package Array

import "fmt"

func minPathSum(grid [][]int) int {
	/*
	Given a m x n grid filled with non-negative numbers,
	find a path from top left to bottom right which minimizes the sum of all numbers along its path.
	Note: You can only move either down or right at any point in time.
	*/

	m := len(grid)
	n := len(grid[0])

	index := 1
	for index < m {
		grid[index][0] = grid[index][0] + grid[index-1][0]
		index += 1
	}

	index = 1
	for index < n {
		grid[0][index] = grid[0][index] + grid[0][index-1]
		index += 1
	}
	fmt.Println(grid)
	left, right := 1, 1
	for left < m {
		right = 1
		for right < n {
			grid[left][right] = min(grid[left-1][right], grid[left][right-1]) + grid[left][right]
			right += 1
		}
		left += 1
	}

	return grid[m-1][n-1]
}


func min(left, right int) int {
	if left < right {
		return left
	} else {
		return right
	}
}