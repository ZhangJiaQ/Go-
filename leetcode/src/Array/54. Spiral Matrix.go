package Array

func spiralOrder(matrix [][]int) []int {

	/*

	Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
	spiral  螺旋的
	*/

	result := make([]int, 0)
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for up <= down &&left <= right {
		if up < len(matrix) {
			temp := left
			for temp <= right {
				result = append(result, matrix[up][temp])
				temp++
			}
			up++
		}
		if right >= 0 {
			temp := up
			for temp <= down {
				result = append(result, matrix[temp][right])
				temp++
			}
			right--
		}
		if down >= up {
			temp := right
			for temp >= left {
				result = append(result, matrix[down][temp])
				temp--
			}
			down--
		}
		if left <= right {
			temp := down
			for temp >= up {
				result = append(result, matrix[temp][left])
				temp--
			}
			left++
		}
	}

	return result

}
