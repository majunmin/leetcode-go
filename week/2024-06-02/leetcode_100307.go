package _024_06_02

// https://leetcode.cn/problems/minimum-number-of-chairs-in-a-waiting-room/
func minimumChairs(s string) int {
	var (
		dp = make([]int, len(s)+1)
	)
	dp[0] = 0
	for i := range s {
		if s[i] == 'E' {
			dp[i+1] = dp[i] + 1
		} else {
			dp[i+1] = max(dp[i]-1, 0)
		}
	}
	return dp[len(dp)-1]
}
