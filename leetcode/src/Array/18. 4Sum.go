package Array

func fourSum(nums []int, target int) [][]int {

	result := make([][]int, 0)
	Quicksort(&nums)
	for index1, _ := range nums {
		if index1 > 0 && nums[index1] == nums[index1 - 1] {
			continue
		}
		index2 := index1 + 1
		for index2 < len(nums) {
			if index2 != index1 + 1 && nums[index2] == nums[index2 - 1] {
				index2 += 1
				continue
			}
			left, right := index2 + 1, len(nums) - 1
			for left < right {
				sum :=  nums[index1] + nums[index2] + nums[left] + nums[right]
				if sum  == target {
					temp := []int{nums[index1] , nums[index2] , nums[left] , nums[right]}
					result = append(result, temp)
					left += 1
					for left < right && nums[left] == nums[left - 1] {
						left += 1
					}
					right -= 1
					for left < right && nums[right] == nums[right + 1] {
						right -= 1
					}
				} else if  sum > target {
					right -= 1
					for left < right && nums[right] == nums[right + 1] {
						right -= 1
					}
				} else {
					left += 1
					for left < right && nums[left] == nums[left - 1] {
						left += 1
					}
				}
			}
			index2 += 1
		}
	}


	return result
}


