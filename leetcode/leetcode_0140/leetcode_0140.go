package leetcode_0140

import "strings"

func wordBreak(s string, wordDict []string) []string {

	// 1， 动态规划 是否有解 (可以参考上一题:139)
	// 2. 题目要求所有可能的结果, 一般就是用回溯算法
	//
	// param check
	if len(s) == 0 || len(wordDict) == 0 {
		return nil
	}
	// 1. dp
	n := len(s)
	dp := make([]bool, len(s)+1)
	// init state, 没有意义， 方便计算
	dp[0] = true
	wordSet := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordSet[word] = true
	}
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			suffix := s[j:i]
			if wordSet[suffix] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	// 2. backtrace 构建路径
	var result []string
	backtrace(&result, n, dp, s, wordSet, []string{})
	return result
}

func backtrace(result *[]string, end int, dp []bool, s string, wordSet map[string]bool, path []string) {
	// terminate
	if end == 0 {
		*result = append(*result, strings.Join(path, " "))
		return
	}
	// for choice in choiceList : doChoice
	for word := range wordSet {
		prevIdx := end - len(word)
		if prevIdx >= 0 && dp[prevIdx] && strings.EqualFold(s[prevIdx:end], word) {
			path = append([]string{word}, path...)
			backtrace(result, prevIdx, dp, s, wordSet, path)
			// revert
			path = path[1:]
		}
	}
}
