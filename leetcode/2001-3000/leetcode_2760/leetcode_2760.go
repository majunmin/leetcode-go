package leetcode_2760

// https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/
func longestAlternatingSubarray(nums []int, threshold int) int {
}

func dpSolution(nums []int, threshold int) int {
	// param check
	if len(nums) == 0 {
		return 0
	}
	var result int
	// 动态规划
	dp := make([]int, len(nums))
	//init state
	if isEven(nums[0]) && nums[0] <= threshold {
		dp[0] = 1
		result = dp[0]
	}

	// 3 2 5 4 4 3  2 1
	// 0 0 1 2 3 1 2  3 4
	//当前位置是偶数， cur = 1 or cur = prev + 1
	//当前位置是奇数， cur = 0 or cur = prev + 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > threshold {
			dp[i] = 0
			continue
		}
		if isDiff(nums[i], nums[i-1]) {
			dp[i] = dp[i-1] + 1
			continue
		}
		if isEven(nums[i]) {
			dp[i] = 1
			continue
		}
	}

	for _, item := range dp {
		if item > result {
			result = item
		}
	}
	return result
}

func isOdd(val int) bool {
	return val&1 == 1
}
func isEven(val int) bool {
	return val&1 == 0
}

// 奇偶不同
func isDiff(val1, val2 int) bool {
	// & 的优先级 > ^
	return (val1^val2)&1 == 1
}
