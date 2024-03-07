package main

import "sort"

// https://leetcode.cn/problems/4sum/
func fourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	result := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := len(nums) - 1
			for {
				if left > j+1 && left < len(nums)-1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				if left >= right {
					break
				}

				if nums[i]+nums[j]+nums[left]+nums[right] == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					continue
				}
				if nums[i]+nums[j]+nums[left]+nums[right] > target {
					right--
					continue
				}
				if nums[i]+nums[j]+nums[left]+nums[right] < target {
					left++
					continue
				}
			}

		}
	}
	return result
}
