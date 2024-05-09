package leetcode_2369

// https://leetcode.cn/problems/check-if-there-is-a-valid-partition-for-the-array/
func validPartition(nums []int) bool {
	// param check
	if len(nums) < 2 {
		return false
	}

	dp := make([]bool, len(nums)+1)
	// init
	dp[0] = true
	dp[1] = false
	// 转移方程
	//dp[i] = dp[i -2] && dp[i] == dp[i-1]
	//      || dp[i -3] && dp[i] == dp[i-1] && dp[i] == dp[i-2] (i > 3)
	//      || dp[i -3] && dp[i]== dp[i-1]+1 && dp[i] == dp[i-2]+2 (i>3)
	for i := 2; i < len(dp); i++ {

		dp[i] = (valid2(nums[i-1], nums[i-2]) && dp[i-2]) ||
			(i > 2 && valid3(nums[i-3], nums[i-2], nums[i-1]) && dp[i-3])
	}
	return dp[len(dp)-1]
}

func valid2(i1, i2 int) bool {
	return i1 == i2
}

func valid3(i1, i2, i3 int) bool {
	return (i1 == i2 && i1 == i3) || (i1+1 == i2 && i2+1 == i3)
}
