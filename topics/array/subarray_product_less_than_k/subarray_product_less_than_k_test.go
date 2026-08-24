package subarrayproductlessthank

import "testing"

func TestNumSubarrayProductLessThanK(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{10, 5, 2, 6}, 100, 8}, // 力扣示例 1
		{[]int{1, 2, 3}, 0, 0},       // 力扣示例 2：k <= 1 → 0
		{[]int{1, 1, 1}, 1, 0},       // k = 1，乘积最小为 1，无解
		{[]int{}, 100, 0},            // 空数组
		{[]int{5}, 10, 1},            // 单元素小于 k
		{[]int{5}, 5, 0},             // 单元素等于 k（不满足严格小于）
	}
	for _, c := range cases {
		if got := numSubarrayProductLessThanK(c.nums, c.k); got != c.want {
			t.Errorf("numSubarrayProductLessThanK(%v, %d) = %d, want %d", c.nums, c.k, got, c.want)
		}
	}
}
