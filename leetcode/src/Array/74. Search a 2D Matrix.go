package Array

func searchMatrix(matrix [][]int, target int) bool {
	// n*m数组中找到target,用二分找
	length := len(matrix)
	if length == 0  {
		return false
	}
	width := len(matrix[0])

	if length == 0 && width == 0 {
		return false
	}

	left, right := 0, length * width - 1
	for left <= right {
		mid := (left + right) / 2
		i, j := mid / width, mid % width
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}