package Dynamic_Programming

import "fmt"

func longestPalindrome(s string) string {
	/*最长回文*/

	if s == ""{
		return ""
	}
	maxLength := 0
	maxIndex := make([]int, 2)
	index := 1
	for index < len(s) {
		if index + 1 < len(s) && s[index-1] == s[index+1] {
			findSubString(&maxLength, &s, index-1, index+1, &maxIndex)
		}
		if s[index-1] == s[index]{
			findSubString(&maxLength, &s, index-1, index, &maxIndex)
		}
		index++
	}
	return s[maxIndex[0]:maxIndex[1]+1]
}


func findSubString(maxLength *int, s *string, left, right int, maxIndex *[]int) {
	for left >= 0 && right < len(*s) {
		if (*s)[left] != (*s)[right] {
			break
		}
		if *maxLength < right - left + 1 {
			*maxLength = right - left
			(*maxIndex)[0] = left
			(*maxIndex)[1] = right
			fmt.Println(left, right)
		}
		left -= 1
		right += 1
	}
}
