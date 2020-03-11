package Array

import sort2 "sort"

func insert(intervals [][]int, newInterval []int) [][]int {

	/*
	Given a set of non-overlapping intervals,
	insert a new interval into the intervals (merge if necessary).
	You may assume that the intervals were initially sorted according to their start times.
	给一个无重叠的区间数组，插入一个新区间
	新区见可能造成原有区间数组的合成

	Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
	Output: [[1,2],[3,10],[12,16]]
	Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].

	Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
	Output: [[1,5],[6,9]]
	*/

	intervals = append(intervals, newInterval)

	if len(intervals) <= 1{
		return intervals
	}

	sort2.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)

	index := 0
	for index < len(intervals){
		start, end := intervals[index][0], intervals[index][1]
		nextIndex := index + 1
		for nextIndex < len(intervals) && end >= intervals[nextIndex][0] {
			if intervals[nextIndex][1] > end {
				end = intervals[nextIndex][1]
			}
			nextIndex++
		}
		result = append(result, []int{start, end})
		index = nextIndex
	}
	return result
}