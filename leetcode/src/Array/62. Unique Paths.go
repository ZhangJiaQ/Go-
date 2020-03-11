package Array

func uniquePaths(m int, n int) int {
	/*
	M*N的棋盘，每次可以选择向右或者向下走一步，求总共能多少种方案
	M=3 N=2  -> return 3
	m=7 n=3  -> return 28
	*/
	// 使用 深度优先遍历
	if m == 1 || n == 1 {
		return 1
	}
	helpArray := make([][]int, m)
	index := 0
	for index < m {
		helpArray[index] = make([]int, n)
		helpArray[index][0] = 1
		index += 1
	}

	index = 0
	for index < n {
		helpArray[0][index] = 1
		index += 1
	}

	left, right := 1, 1
	for left < m {
		right = 1
		for right < n {
			helpArray[left][right] = helpArray[left - 1][right] + helpArray[left][right - 1]
			right += 1
		}
		left += 1
	}


	return helpArray[m-1][n-1]
}