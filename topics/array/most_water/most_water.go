package mostwater

func maxArea(height []int) (ans int) {
	left := 0
	right := len(height) - 1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		ans = max(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return
}
