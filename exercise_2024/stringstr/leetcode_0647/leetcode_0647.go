package leetcode_0647

import "fmt"

// https://leetcode.cn/problems/palindromic-substrings/
func countSubstrings(s string) int {
	return solution3(s)
}

// 动态规划解法
func solution3(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	// dp
	// 状态定义
	//dp[i][j] 字符串s[i:j] 是否是回文子串
	//dp[i][j] = dp[i-1][j+1] && s[i] == s[j] // i-j > 2
	//dp[i][j] = s[i] == s[j] // i-j = 1
	var result int
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				result++
			}
		}
	}
	return result
}

// 中心扩展法
func solution2(s string) int {
	var result int
	for i := 0; i < len(s)*2; i++ {
		left, right := i>>1, (i+1)>>1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			fmt.Println(s[left : right+1])
			result++
			left--
			right++
		}
	}
	return result
}

// 暴力枚举法
func baolijiefa1(s string) int {
	//暴力解法
	var result int
	for i := 1; i <= len(s); i++ {
		for j := 0; i+j <= len(s); j++ {
			if isPalindromic(s[j : j+i]) {
				result++
			}
		}
	}
	return result
}

func isPalindromic(s string) bool {
	fmt.Println(s)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
