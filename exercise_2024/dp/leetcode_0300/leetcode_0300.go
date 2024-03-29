package leetcode_0300

// https://leetcode.cn/problems/maximum-product-subarray/?envType=study-plan-v2&envId=top-100-liked
func lengthOfLIS(nums []int) int {

	return dpSolution2(nums)
}

// 动态规划 + 二分搜索
func dpSolution2(nums []int) int {
	tail := make([]int, len(nums)+1)
	// tails[i] 表示长度为 i 的 子序列 的最小tail值
	res := 1
	for _, num := range nums {
		l, r := 1, res
		// 二分查找法 找到左边第一个 <= target 的值
		for l < r {
			mid := l + (r-l)>>1
			if num <= tail[mid] {
				r = mid
			} else {
				l = mid + 1
			}
		}
		tail[l] = num
		if res == l {
			res++
		}
	}
	return res - 1
}

func dpSolution(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	// param check
	dp := make([]int, len(nums))
	// init state
	var result int
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}
	return result
}
