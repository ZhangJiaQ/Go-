package lintcode

/**
 * @param A: sorted integer array A
 * @param B: sorted integer array B
 * @return: A new sorted integer array
 */
func mergeSortedArray (A []int, B []int) []int {
	// write your code here

	aLen := len(A)
	bLen := len(B)

	C := make([]int, 0)

	var aContent, bContent int

	for aContent < aLen && bContent < bLen {
		if A[aContent] < B[bContent] {
			C = append(C, A[aContent])
			aContent += 1
		} else {
			C = append(C, B[bContent])
			bContent += 1
		}
	}

	for aContent < aLen {
		C = append(C, A[aContent])
		aContent += 1
	}

	for bContent < bLen {
		C = append(C, B[bContent])
		bContent += 1
	}

	return C

}


