package Dynamic_Programming

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2{
		return 2
	}
	result := make([]int, n)
	result[0] = 1
	result[1] = 2
	index := 2
	for index < n {
		result[index] = result[index-1] + result[index-2]
		index += 1
	}
	return result[n-1]

}

