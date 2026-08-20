package threesum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	cases := []struct {
		nums []int
		want [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}}, // 标准用例
		{[]int{0, 1, 1}, nil},                                          // 无解
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},                           // 全零
		{[]int{}, nil},                                                 // 空数组
		{[]int{1, 2}, nil},                                             // 不足 3 个
	}
	for _, c := range cases {
		if got := ThreeSum(c.nums); !reflect.DeepEqual(got, c.want) {
			t.Errorf("ThreeSum(%v) = %v, want %v", c.nums, got, c.want)
		}
	}
}
