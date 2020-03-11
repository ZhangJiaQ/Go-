package Array

func twoSum(nums []int, target int) []int {
	// 使用Map作为存储，找target - nownum 是否存在于Map中且不等于自身即可

	arrayMap := make(map[int][]int, 0)

	for index, v := range nums {
		arrayMap[v] = append(arrayMap[v], index)
	}

	for index, value := range arrayMap{
		otherSum := target - index
		if otherSum == index {
			if len(value) > 1{
				return []int{value[0], value[1]}
			}
		} else {
			otherValue, ok := arrayMap[otherSum]
			if ok{
				if value[0] < otherValue[0]{
					return []int{value[0], otherValue[0]}
				} else {
					return []int{otherValue[0], value[0]}
				}
			}
		}

	}
	return []int{0,0}
}