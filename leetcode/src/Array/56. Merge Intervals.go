package Array

import (
	sort2 "sort"
	"fmt"
)

func merge(intervals [][]int) [][]int {
	/*
	合并区间
	*/
	if len(intervals) <= 1 {
		return intervals
	}

	// 排序
	sort2.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)

	index := 0
	for index < len(intervals) {
		start, end := intervals[index][0], intervals[index][1]
		nextIndex := index + 1
		for nextIndex < len(intervals) && intervals[nextIndex][0] <= end {
			// 有交集，从交集从取出最大end
			if intervals[nextIndex][1] > end {
				end = intervals[nextIndex][1]
			}
			nextIndex++
		}
		// 无交集，把当前start,end添加至结果数组
		fmt.Println(start, end)
		result = append(result, []int{start, end})
		index = nextIndex
	}

	return result

}
