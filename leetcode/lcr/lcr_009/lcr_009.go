package lcr_009

// https://leetcode.cn/problems/ZVAVXX/
func numSubarrayProductLessThanK(nums []int, k int) int {
	//滑动窗口
	// param check
	var (
		n          = len(nums)
		l, r       int
		result     int
		windowProd = 1
	)
	for ; r < n; r++ {
		// invalid: 子数组乘积 >=k
		windowProd *= nums[r]
		for ; l < r && windowProd >= k; l++ {
			// shrink window
			windowProd /= nums[l]
		}
		result += r - l + 1
	}
	return result
}

//100
// 10 5 2 6
// 1 10 5 2 6 20 4 12
// 2
// 3 30 15 6
