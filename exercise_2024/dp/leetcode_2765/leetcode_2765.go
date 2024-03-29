package leetcode_2765

func alternatingSubarray(nums []int) int {
	return solution2(nums)
}

// 分组循环
// https://leetcode.cn/problems/longest-alternating-subarray/solutions/2615916/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-r57bz/
func solution1(nums []int) int {
	// param check
	result := 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			continue
		}
		cnt := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[j-2] {
				cnt += 1
			}
		}
		result = max(result, cnt)
	}
	if result == 1 {
		return -1
	}
	return result
}

// 动态规划解法 可以优化空间
func solution2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			if dp[i]&1 == 1 {
				dp[i] += 1
			} else {
				dp[i] = dp[i-1]
			}
			continue
		}
		if nums[i]-nums[i-1] == -1 {
			if dp[i]&1 == 0 {
				dp[i] += 1
			} else {
				dp[i] = 1
			}
			continue
		}
		dp[i] = 1
	}
	result := 1
	for i := 0; i < len(dp); i++ {
		result = max(result, dp[i])
	}
	return result
}
