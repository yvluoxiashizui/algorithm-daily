package threesum

import "slices"

func ThreeSum(nums []int) (ans [][]int) {
	if len(nums) < 3 {
		return
	}
	slices.Sort(nums)
	n := len(nums)
	//遍历，将三数之和转化为两数之和=-x
	for i, x := range nums[:n-2] {
		//去重
		if i > 0 && x == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		//剪枝优化，最小数之和>0，后面只会更大，直接不用看了
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
			//最大数之和<0，但nums[i]还能变大，只continue跳过当前循环
		} else if nums[i]+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		for j < k {
			s := nums[i] + nums[j] + nums[k]
			if s < 0 {
				j++
			} else if s > 0 {
				k--
			} else {
				ans = append(ans, []int{x, nums[j], nums[k]})
				//去重
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return
}
