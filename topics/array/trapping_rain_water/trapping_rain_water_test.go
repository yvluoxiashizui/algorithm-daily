package trapping_rain_water

import "testing"

func TestTrap(t *testing.T) {
	cases := []struct {
		height []int
		want   int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6}, // 力扣示例 1
		{[]int{4, 2, 0, 3, 2, 5}, 9},                    // 力扣示例 2
		{[]int{}, 0},                                    // 空数组（加了保护才不 panic）
		{[]int{5}, 0},                                   // 单元素
		{[]int{3, 1, 2}, 1},                             // 简单情况：中间洼地存 1
	}
	for _, c := range cases {
		if got := trap(c.height); got != c.want {
			t.Errorf("trap(%v) = %d, want %d", c.height, got, c.want)
		}
	}
}
