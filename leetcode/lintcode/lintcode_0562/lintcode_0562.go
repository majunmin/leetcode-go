package lintcode_0562

func BackPackIV(nums []int, target int) int {
	// write your code here
	// param check
	if len(nums) == 0 {
		return 0
	}
	//
	dp := make([]int, target+1)
	// init state
	dp[0] = 1
	for _, num := range nums {
		for j := num; j <= target; j++ {
			dp[j] += dp[j-num]
		}
	}
	return dp[target]
}
