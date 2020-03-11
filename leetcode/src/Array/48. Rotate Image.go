package Array

func rotate(matrix [][]int)  {
	/*
	 1 2 3
	 4 5 6
	 7 8 9

	 - >

	 7 4 1
	 8 5 2
	 9 6 3


	 1 2  3  4
	 5 6  7  8
	 9 10 11 12
	13 14 15 16
	 保证 N * N的数组
	 要求原地转

	*/



	startLeft := 0

	length := len(matrix)

	for startLeft < length / 2 {
		col := startLeft
		row := startLeft
		// temp := matrix[col][row]
		// 移动步数
		step := length - (startLeft * 2) - 1
		for {
			if col == startLeft && row == startLeft + 1 {
				break
			}
			for col < length - startLeft {
				// 移动
				diff := col - step
				if diff < startLeft {
					matrix[col][row] = matrix[startLeft][row - diff + startLeft]
				} else {
					matrix[col][row] = matrix[col-step][row]
				}
				col ++
			}
			col --
			for row < length - startLeft{
				row ++
			}
			row --
			for col > startLeft {
				col --
			}
			for row > startLeft {
				row --
			}
			// 更改temp
			// temp =  matrix[col][row]



			// 调整row col位置

			if col == startLeft && row == startLeft + 1 {
				break
			}
		}

		startLeft += 1
	}


}