package leetcode2529

import "sort"

func maximumCount(nums []int) int {
	neg := sort.SearchInts(nums, 0)
	// 第一个 > 0 的位置，等价于第一个 >= 1 的位置
	pos := len(nums) - sort.SearchInts(nums, 1)
	return max(neg, pos)
}
