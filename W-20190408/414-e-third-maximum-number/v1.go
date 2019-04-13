package leetcode

import "sort"

func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	c := 0
	l := len(nums)
	m := nums[l-1]
	for i := l - 1; i >= 0; i-- {
		if nums[i] < m {
			m = nums[i]
			c++
			if c == 2 {
				return m
			}
		}
	}
	return nums[l-1]
}
