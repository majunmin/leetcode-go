package leetcode_1617

// https://leetcode.cn/problems/minimum-number-of-removals-to-make-mountain-array/
func minimumMountainRemovals(nums []int) int {
	// 1. 获取 以每一个点作为山脉的顶点的最长递增子序列长度 pre[]
	// 2. 获取 以每一个点作为山脉的顶点的最长递减子序列长度 suf[]
	pre := getLISArray(nums)
	reverse(nums)
	suf := getLISArray(nums)
	reverse(suf)

	var (
		n      = len(nums)
		result = n
	)
	for i := 1; i < len(pre)-1; i++ {
		if pre[i] == 1 || suf[i] == 1 {
			continue
		}
		result = min(result, n-pre[i]-suf[i]+1)
	}
	return result
}

// 最长连续递增子序列
func getLISArray(nums []int) []int {
	var (
		n  = len(nums)
		dp = make([]int, n)
	)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
