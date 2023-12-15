package leetcode_0238

// https://leetcode.cn/problems/product-of-array-except-self/?envType=study-plan-v2&envId=top-100-liked
func productExceptSelf(nums []int) []int {
	return solution2(nums)
}

// 空间优化的解法
func solution2(nums []int) []int {
	// param check
	if len(nums) < 2 {
		return nums
	}
	n := len(nums)
	result := make([]int, n)
	result[0] = 1
	for i := 1; i < n; i++ {
		// result 相当于 left 数组
		result[i] = result[i-1] * nums[i-1]
	}
	rSum := 1
	for i := len(nums) - 2; i >= 0; i-- {
		rSum = rSum * nums[i+1]
		result[i] = result[i] * rSum
	}

	return result
}

func solution(nums []int) []int {
	// result[i] = left[i] * right[i]
	// left[i]表示 区间[0, i-1] 的乘积，(i > 0)
	left, right := make([]int, len(nums)), make([]int, len(nums))
	// init
	left[0], right[len(right)-1] = 1, 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
		right[len(nums)-i-1] = right[len(nums)-i] * nums[len(nums)-i]
	}
	result := make([]int, len(nums))
	for i := 0; i < len(result); i++ {
		result[i] = left[i] * right[i]
	}
	return result
}
