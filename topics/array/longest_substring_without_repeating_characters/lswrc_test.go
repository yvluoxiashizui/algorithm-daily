package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3}, // 力扣示例 1
		{"bbbbb", 1},    // 力扣示例 2
		{"pwwkew", 3},   // 力扣示例 3
		{"", 0},         // 空字符串
		{"au", 2},       // 简单两个字符
		{"dvdf", 3},     // 陷阱：答案跨过中间的重复字符
	}
	for _, c := range cases {
		if got := lengthOfLongestSubstring(c.s); got != c.want {
			t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", c.s, got, c.want)
		}
	}
}
