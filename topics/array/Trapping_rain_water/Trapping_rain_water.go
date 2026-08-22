package trapping_rain_water

func trap(height []int) (ans int) {
	n := len(height)
	if n == 0 {
		return 0
	}
	pro_max := make([]int, n)
	suf_max := make([]int, n)
	pro_max[0] = height[0]
	suf_max[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		pro_max[i] = max(pro_max[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		suf_max[i] = max(suf_max[i+1], height[i])
	}
	for i := 0; i < n; i++ {
		hei := min(pro_max[i], suf_max[i]) - height[i]
		if hei > 0 {
			ans += hei
		}
	}
	return ans
}
