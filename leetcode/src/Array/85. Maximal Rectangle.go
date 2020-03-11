package Array


func maximalRectangle(matrix [][]byte) int {

	/*
	Given a 2D binary matrix filled with 0's and 1's,
	find the largest rectangle containing only 1's and return its area.
	*/
	if len(matrix) == 0 {
		return 0
	}

	// 辅助数组使用
	for i1, v1 := range matrix {
		for i2, v2 := range v1 {
			if v2 == '1' {
				// 确定长 , 确定长的时候更新辅助数组

				// 确定高
			}
		}
	}
	return 0

}