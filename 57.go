package main

// https://leetcode.cn/problems/insert-interval/description/
func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	left := newInterval[0]
	right := newInterval[1]
	for _, it := range intervals {
		if it[0] > right {
			result = append(result, []int{left, right})
			left = it[0]
			right = it[1]
			continue
		}
		if it[1] < left {
			result = append(result, []int{it[0], it[1]})
			continue
		}
		left = min(left, it[0])
		right = max(right, it[1])

	}
	result = append(result, []int{left, right})
	return result
}
