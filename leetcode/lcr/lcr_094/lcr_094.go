package lcr_094

// https://leetcode.cn/problems/omKAoA/
func minCut(s string) int {
	// 动态规划问题
	var (
		dp = make([]int, len(s)+1)
	)
	// init state
	dp[0] = 0
	for j := 1; j <= len(s); j++ {
		for i := j - 1; i >= 0; i-- {
			if isPalindrome(s[i:j]) {
				dp[j] = min(dp[j], dp[i]+1)
			}
		}
	}
	return dp[len(dp)] - 1
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
