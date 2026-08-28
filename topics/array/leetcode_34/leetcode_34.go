package leetcode34

func lowerBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func searchRange(nums []int, target int) []int {
	start := lowerBound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := lowerBound(nums, target+1) - 1
	return []int{start, end}
}
