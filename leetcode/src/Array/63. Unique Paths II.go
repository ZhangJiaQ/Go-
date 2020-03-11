package Array

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 还是走路那道题，不过中间有挡路的了

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if m ==1  || n == 1 {
		for _, v1 := range obstacleGrid {
			for _, v2 := range v1 {
				if v2 == 1{
					return 0
				}
			}
		}

		return 1
	}

	if  obstacleGrid[m-1][n-1] == 1 {
		return 0
	}


	for index1, v1 := range obstacleGrid{
		for index2, v2 := range v1 {
			if v2 == 1 {
				obstacleGrid[index1][index2] = -1
			}
		}
	}

	index := 0
	for index < m {
		if obstacleGrid[index][0] == -1 {
			break
		}
		obstacleGrid[index][0] = 1
		index += 1
	}


	index = 0
	for index < n {
		if obstacleGrid[0][index] == -1 {
			break
		}
		obstacleGrid[0][index] = 1
		index += 1
	}

	left, right := 1, 1
	for left < m {
		right = 1
		for right < n {
			if obstacleGrid[left][right] == -1 {
				right += 1
				continue
			}
			if obstacleGrid[left-1][right] == -1 && obstacleGrid[left][right - 1] == -1   {
				obstacleGrid[left][right] = 0
			} else if obstacleGrid[left][right - 1] == -1 {
				obstacleGrid[left][right] = obstacleGrid[left - 1][right]
			} else if obstacleGrid[left-1][right] == -1 {
				obstacleGrid[left][right] = obstacleGrid[left][right - 1]
			} else {
				obstacleGrid[left][right] = obstacleGrid[left][right - 1] + obstacleGrid[left - 1][right]
			}
			right += 1
		}
		left += 1
	}

	return obstacleGrid[m-1][n-1]
}