package Array

func generateMatrix(n int) [][]int {

	/*
	给一个n
	产生一个N * N的且顺时针旋转一次的数组
	*/
	// 构建一个N*N的数组
	result := make([][]int, n)
	index := 0
	for index < n {
		result[index] = make([]int, n)
		index++
	}
	// 往里填数字
	number := 1

	left, right, up, down := 0, n - 1, 0, n - 1
	for number < n * n {
		if left <= right {
			temp := left
			for temp <= right {
				result[up][temp] = number
				number++
				temp++
			}
			up += 1
		}
		if up <= down {
			temp := up
			for temp <= down {
				result[temp][right] = number
				number++
				temp++
			}
			right -= 1
		}
		if right >= left {
			temp := right
			for temp >= left {
				result[down][temp] = number
				temp--
				number++
			}
			down--
		}
		if down >= up{
			temp := down
			for temp >= up {
				result[temp][left] = number
				number++
				temp--
			}
			left++
		}
	}

	if n % 2 == 1 {
		result[n/2][n/2] = n * n
	}

	return result
}
