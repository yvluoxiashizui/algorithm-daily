package minimum_size_subarray_sum

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	cases := []struct {
		target int
		nums   []int
		want   int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},                        // 力扣示例 1
		{4, []int{1, 4, 4}, 1},                                 // 力扣示例 2
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},                 // 力扣示例 3：总和都不够 → 0
		{15, []int{1, 2, 3, 4, 5}, 5},                          // 全部元素刚好达标
		{1, []int{1}, 1},                                       // 单元素正好
		{100, []int{}, 0},                                      // 空数组
	}
	for _, c := range cases {
		if got := minSubArrayLen(c.target, c.nums); got != c.want {
			t.Errorf("minSubArrayLen(%d, %v) = %d, want %d", c.target, c.nums, got, c.want)
		}
	}
}
