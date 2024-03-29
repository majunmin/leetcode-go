package leetcode_0238

// https://leetcode.cn/problems/product-of-array-except-self/?envType=study-plan-v2&envId=top-100-liked
func productExceptSelf(nums []int) []int {
	// param check
	n := len(nums)
	if n == 0 {
		return nil
	}
	lp, rp := make([]int, n), make([]int, n)
	lp[0], rp[n-1] = 1, 1
	for i := 1; i < n; i++ {
		lp[i] = lp[i-1] * nums[i-1]
		rp[n-1-i] = rp[n-i] * nums[n-i]
	}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = lp[i] * rp[i]
	}
	return result
}
