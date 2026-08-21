package mostwater

import "testing"

func TestMaxArea(t *testing.T) {
	cases := []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49}, // 力扣示例
		{[]int{1, 1}, 1},                       // 两个柱子
		{[]int{1, 2, 1}, 2},                    // 中间高两边低
		{[]int{1, 3, 2, 5, 25, 24, 5}, 24},     // 陷阱：最大面积不一定是最高两柱之间
	}
	for _, c := range cases {
		if got := maxArea(c.height); got != c.want {
			t.Errorf("maxArea(%v) = %d, want %d", c.height, got, c.want)
		}
	}
}
