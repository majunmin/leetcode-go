package leetcode_1143

// https://leetcode.cn/problems/longest-common-subsequence/
func longestCommonSubsequence(text1 string, text2 string) int {
	l1, l2 := len(text1), len(text2)
	if l1 == 0 || l2 == 0 {
		return 0
	}

	// init dp state
	dp := make([][]int, 0, l1+1)
	for i := 0; i < l1; i++ {
		dp = append(dp, make([]int, l2+1))
	}

	//
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]

}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
