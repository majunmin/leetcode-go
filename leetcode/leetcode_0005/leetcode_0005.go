package leetcode_0005

// https://leetcode.cn/problems/longest-palindromic-substring/?envType=study-plan-v2&envId=top-100-liked
func longestPalindrome(s string) string {
	return dpSolution(s)
}

func dpSolution(s string) string {
	// param check
	n := len(s)
	if n < 2 {
		return s
	}
	// 状态定义
	// 闭区间： dp[i][j]: s[i:j+1] 是否是一个 回文串
	// dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
	// init state
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	// 动态规划
	result := string(s[0])
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] != s[j] {
				dp[i][j] = false
				continue
			}
			// s[i] == s[j]
			if j-i < 2 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
			}
			if dp[i][j] && j-i+1 > len(result) {
				result = s[i : j+1]
			}
		}
	}
	return result
}

// 中心扩展算法,枚举中心点, 计算最长回文串
func solution1(s string) string {
	// param check
	if len(s) == 0 {
		return ""
	}
	var result string
	for i := 0; i < len(s); i++ {
		str1 := extensions(s, i, i)
		if len(str1) > len(result) {
			result = str1
		}
		str2 := extensions(s, i, i+1)
		if len(str2) > len(result) {
			result = str2
		}
	}
	return result
}

func extensions(s string, l int, r int) string {
	length := len(s)
	for l >= 0 && r < length {
		if s[l] != s[r] {
			break
		}
		l--
		r++
	}
	// ab
	if r-l == 1 {
		return ""
	}
	return s[l+1 : r-1]
}
