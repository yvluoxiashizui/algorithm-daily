package twosumii

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}}, // 基础：答案在两端
		{[]int{2, 3, 4}, 6, []int{1, 3}},      // 答案不相邻
		{[]int{-1, 0}, -1, []int{1, 2}},       // 含负数
		{[]int{1, 2, 3}, 99, []int{}},         // 无解
	}
	for _, c := range cases {
		if got := TwoSum(c.numbers, c.target); !reflect.DeepEqual(got, c.want) {
			t.Errorf("TwoSum(%v, %d) = %v, want %v", c.numbers, c.target, got, c.want)
		}
	}
}
