package Array

func exist(board [][]byte, word string) bool {
	/*
	遍历board，找到board与word第一个相等的字符 开始业务逻辑
	业务逻辑
	1. 找到上下左右且符合下一个字符的 如果符合 （最多四个）加入栈
	2. 从栈中取出头部元素，继续搜索上下左右但不包括来源位置的值，找到后（最多三个）加入栈
	3.
	*/

	helpArray := make([][]bool, len(board))
	for index := range helpArray {
		helpArray[index] = make([]bool, len(board[0]))
	}

	for indexRow, valueRow := range board {
		for indexCol, valueCol := range valueRow {
			if valueCol == word[0]{
				// 调用Search方法
				if searchDfs(&board, indexRow, indexCol, 0, &helpArray, word) {
					return true
				}
			}
		}
	}
	return false
}


func searchDfs(board *[][]byte, indexRow, indexCol, wordIndex int, helpArray *[][]bool, word string) bool {
	if wordIndex == len(word) {
		return true
	}

	// 越界检查
	if indexRow < 0 || indexCol < 0 || indexRow > len(*board) - 1 || indexCol > len((*board)[0]) - 1 {
		return false
	}

	// 检查过的不在检查
	if (*helpArray)[indexRow][indexCol] == true {
		return false
	}

	if (*board)[indexRow][indexCol] != word[wordIndex] {
		return false
	}

	(*helpArray)[indexRow][indexCol] = true
	result := searchDfs(board, indexRow-1, indexCol, wordIndex+1, helpArray, word) || searchDfs(board, indexRow+1, indexCol, wordIndex+1, helpArray, word) || searchDfs(board, indexRow, indexCol+1, wordIndex+1, helpArray, word) || searchDfs(board, indexRow, indexCol-1, wordIndex+1, helpArray, word)
	(*helpArray)[indexRow][indexCol] = false

	return result

}