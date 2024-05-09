package leetcode_0131

// https://leetcode.cn/problems/palindrome-partitioning/?envType=study-plan-v2&envId=top-100-liked
func partition(s string) [][]string {
	return dpAndBacktraceSolution(s)
}

// 参考  leetcode-0005:https://leetcode.cn/problems/longest-palindromic-substring/?envType=study-plan-v2&envId=top-100-liked
// 1. 先利用动态规划计算出所有的回文串,
// 2. 在利用回溯的方式,优化 isPalindrome 方法
func dpAndBacktraceSolution(s string) [][]string {
	// 状态定义: dp[i][j]: s[i][j+1] 是否是一个回文串
	// init state
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	// 状态转移
	// dp[i][j] = dp[i+1][j-1] && (s[i] == s[j])
	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] != s[j] {
				continue
			}
			dp[i][j] = (j-i == 1) || dp[i+1][j-1]
		}
	}
	var result [][]string
	// 回溯算法
	backtrace2(s, 0, dp, &result, []string{})
	return result
}

func backtrace2(s string, begin int, dp [][]bool, result *[][]string, path []string) {
	// terminate
	if begin == len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}
	// for choice in choiceList
	for i := begin; i < len(s); i++ {
		if !dp[begin][i] {
			continue
		}
		path = append(path, s[begin:i+1])
		backtrace2(s, i+1, dp, result, path)
		path = path[:len(path)-1]
	}
}

// 回溯的方式解决
func backtraceSolution(s string) [][]string {
	// param check
	if len(s) == 0 {
		return nil
	}
	//
	var result [][]string
	backtrace(s, 0, &result, []string{})
	return result
}

func backtrace(s string, begin int, result *[][]string, path []string) {
	// terminate
	if begin == len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}
	// for choice in choiceList
	for i := begin; i < len(s); i++ {
		if !isPalindrome(s, begin, i) {
			continue
		}
		path = append(path, s[begin:i+1])
		backtrace(s, i+1, result, path)
		// revert
		path = path[:len(path)-1]
	}
}

func isPalindrome(s string, l int, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
