package Array
func setZeroes(matrix [][]int)  {

	/*
	Given a m x n matrix,
	if an element is 0, set its entire row and column to 0. Do it in-place.
	*/
	allNoneZeroColoum :=  -1

	// 找到一个全不为0的行
	for i1, v1 := range matrix {
		endFlag := false
		for _, v2 := range v1 {
			if v2 == 0 {
				endFlag = true
				break
			}
		}
		if !endFlag {
			allNoneZeroColoum = i1
		}
	}
	if allNoneZeroColoum == -1 {
		// 全0即可
		for i1  := range matrix{
			for i2 := range matrix[i1] {
				matrix[i1][i2] = 0
			}
		}
	} else {
		for i1  := range matrix{
			for i2 := range matrix[i1] {
				if matrix[i1][i2] == 0 && matrix[allNoneZeroColoum][i2] != 0 {
					matrix[allNoneZeroColoum][i2] = 0
				}
			}
		}
		for i1  := range matrix{
			if i1 == allNoneZeroColoum {
				continue
			}
			zeroFindFlag := false
			for i2 := range matrix[i1] {
				if matrix[i1][i2] == 0{
					zeroFindFlag = true
				}
				if matrix[allNoneZeroColoum][i2] == 0 {
					matrix[i1][i2] = 0
				}
			}
			if zeroFindFlag {
				for i2 := range matrix[i1] {
					matrix[i1][i2] = 0
				}
			}
		}
	}
}
