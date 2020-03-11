package Array

func plusOne(digits []int) []int {

	/*
	Given a non-empty array of digits representing a non-negative integer, plus one to the integer.
	The digits are stored such that the most significant digit is at the head of the list,
	and each element in the array contain a single digit.
	You may assume the integer does not contain any leading zero, except the number 0 itself.
	*/

	digits[len(digits)-1] = digits[len(digits)-1] + 1

	index := len(digits) - 1

	carryFlag := false

	for index >= 0 {
		if index != 0 {
			if digits[index] == 10 {
				digits[index] = 0
				digits[index - 1] = digits[index - 1] + 1
			}
		} else {
			if digits[index] == 10 {
				carryFlag = true
				digits[index] = 0
			}
		}
		index -= 1
	}

	if carryFlag {
		digits = append([]int{1}, digits...)
	}

	return digits

}