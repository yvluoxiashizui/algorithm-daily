package leetcode34

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},  // 力扣示例 1
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}}, // 力扣示例 2：不存在
		{[]int{}, 0, []int{-1, -1}},                  // 空数组
		{[]int{1}, 1, []int{0, 0}},                   // 单元素
		{[]int{2, 2}, 2, []int{0, 1}},                // 全相同
		{[]int{1, 2, 3}, 1, []int{0, 0}},             // 目标在开头
		{[]int{1, 2, 3}, 3, []int{2, 2}},             // 目标在结尾
	}
	for _, c := range cases {
		if got := searchRange(c.nums, c.target); !reflect.DeepEqual(got, c.want) {
			t.Errorf("searchRange(%v, %d) = %v, want %v", c.nums, c.target, got, c.want)
		}
	}
}
