package leetcode2529

import "testing"

func TestMaximumCount(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{-2, -1, -1, 1, 2, 3}, 3},    // 力扣示例 1
		{[]int{-3, -2, -1, 0, 0, 1, 2}, 3}, // 力扣示例 2：含 0
		{[]int{5, 20, 66, 1314}, 4},        // 力扣示例 3：全正
		{[]int{-2, -1, -1}, 3},             // 全负
		{[]int{0, 0, 0}, 0},                // 全 0
		{[]int{}, 0},                       // 空数组
		{[]int{-1, -1, 0, 0, 1, 1}, 2},     // 正负相等
	}
	for _, c := range cases {
		if got := maximumCount(c.nums); got != c.want {
			t.Errorf("maximumCount(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}
