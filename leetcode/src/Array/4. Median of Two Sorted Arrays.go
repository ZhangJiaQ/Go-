package Array

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	twoArrayLen := len(nums1) + len(nums2)

	if twoArrayLen % 2 == 0 {
		// 双数情况下中位数位两个数中间两个之和除以2
		needFindIndex := twoArrayLen / 2
		index1 := 0
		index2 := 0
		number1 := -1
		number2 := -1
		for (index1 + index2) <= needFindIndex {

			if index1 < len(nums1) && index2 < len(nums2){
				// nums1 和 nums2 都还有数据
				if nums1[index1] < nums2[index2] {
					number2 = number1
					number1 = nums1[index1]
					index1 += 1
				} else {
					number2 = number1
					number1 = nums2[index2]
					index2 += 1
				}
			} else if index1 >= len(nums1) && index2 < len(nums2) {
				// nums1 空
				number2 = number1
				number1 = nums2[index2]
				index2 += 1
			}  else if index1 < len(nums1) && index2 >= len(nums2) {
				// nums2 空
				number2 = number1
				number1 = nums1[index1]
				index1 += 1
			}
		}
		return (float64(number1) + float64(number2)) / 2
	} else {
		// 单数情况下找到第 twoArrayLen / 2大的数即可
		needFindIndex := twoArrayLen / 2
		index1 := 0
		index2 := 0
		number := -1
		for (index1 + index2) <= needFindIndex {

 			if index1 < len(nums1) && index2 < len(nums2){
				// nums1 和 nums2 都还有数据
				if nums1[index1] < nums2[index2] {
					number = nums1[index1]
					index1 += 1
				} else {
					number = nums2[index2]
					index2 += 1
				}
			} else if index1 >= len(nums1) && index2 < len(nums2) {
				// nums1 空
				number = nums2[index2]
				index2 += 1
			}  else if index1 < len(nums1) && index2 >= len(nums2) {
				// nums2 空
				number = nums1[index1]
				index1 += 1
			}
		}
		return float64(number)

	}
	return -1.0
}