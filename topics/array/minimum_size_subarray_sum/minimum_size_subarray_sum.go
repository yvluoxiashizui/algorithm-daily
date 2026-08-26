package minimum_size_subarray_sum

func minSubArrayLen(target int, nums []int) (ans int) {
	n := len(nums)
	left := 0
	s := 0
	ans = n + 1
	for right, x := range nums {
		s += x
		for s >= target {
			ans = min(ans, right-left+1)
			s -= nums[left]
			left += 1
		}
	}
	if ans == n+1 {
		return 0
	}
	return ans
}