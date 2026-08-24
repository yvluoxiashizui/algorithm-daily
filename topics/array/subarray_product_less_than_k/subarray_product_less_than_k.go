package subarrayproductlessthank

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	left := 0
	s := 1
	if k <= 1 {
		return 0
	}
	for right, x := range nums {
		s *= x
		for s >= k {
			s = s / nums[left]
			left += 1
		}
		ans += right - left + 1
	}
	return ans
}
