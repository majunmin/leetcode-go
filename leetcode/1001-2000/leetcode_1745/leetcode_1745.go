package leetcode_1745

// https://leetcode.cn/problems/palindrome-partitioning-iv/
func checkPartitioning(s string) bool {
	var (
		n           = len(s)
		isPalindrom = make([][]bool, n)
	)
	// 状态转移方程.
	// dp[i][j] = s[i] == s[j] && (j-i <=2 || dp[i+1][j-1])
	for i := n - 1; i >= 0; i-- {
		isPalindrom[i] = make([]bool, n)
		for j := i; j < n; j++ {
			isPalindrom[i][j] = s[i] == s[j] && (j-i <= 2 || isPalindrom[i+1][j-1])
		}
	}
	// 枚举三个字符串
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n; j++ {
			if isPalindrom[0][i] && isPalindrom[i+1][j-1] && isPalindrom[j][n-1] {
				return true
			}
		}
	}
	return false
}
