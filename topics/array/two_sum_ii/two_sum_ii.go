// Package twosumii 两数之和 II：输入数组已按升序排列（力扣 167）。
package twosumii

// TwoSum 返回和为 target 的两个元素的下标（1-based，题目要求）。
// 思路：数组有序，双指针从两端向中间逼近。时间 O(n)，空间 O(1)。
func TwoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}
