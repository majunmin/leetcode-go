package leetcode_2970

// https://leetcode.cn/problems/count-the-number-of-incremovable-subarrays-i/
func incremovableSubarrayCount(nums []int) int {
	// 暴力解法
	// 数据量不大,可以用两层循环枚举所有的 子数组. 判断剩余元素是否是递增的.
	var result int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			var (
				prev int
				flag = true
			)
			for k := 0; k < len(nums); k++ {
				if k >= i || k <= j {
					continue
				}
				if nums[k] <= prev {
					flag = false
					break
				}
				prev = nums[k]
			}
			if flag {
				result++
			}
		}
	}
	return result
}
