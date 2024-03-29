package leetcode_0139

import "strings"

func wordBreak(s string, wordDict []string) bool {
	// param check
	if len(s) == 0 || len(wordDict) == 0 {
		return false
	}
	return solution3(s, wordDict)
}

// 动态规划,可以抽象为 01背包问题
func solution3(s string, dict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range dict {
		wordSet[word] = true
	}
	// process
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := i; j >= 0; j-- {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

// bfs solution
func solution2(s string, dict []string) bool {
	queue := make([]int, 0)
	queue = append(queue, 0)
	wordSet := make(map[string]bool)
	for _, word := range dict {
		wordSet[word] = true
	}
	visited := make(map[int]bool)
	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		if visited[start] {
			continue
		}
		visited[start] = true
		for j := start + 1; j <= len(s); j++ {
			prefix := s[start:j]
			if wordSet[prefix] {
				if j == len(s) {
					return true
				} else {
					queue = append(queue, j)
				}
			}
		}
	}
	return false
}

// 1. 添加记忆化的  dfs方法
func solution1(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, item := range wordDict {
		wordSet[item] = true
	}
	// memo[i] 以 s[i:] 是否可以由 wordDict 拼接成
	memo := make(map[int]bool, len(s)) // 记忆化
	return dfs(s, 0, memo, wordSet)
}

func dfs(s string, begin int, memo map[int]bool, wordSet map[string]bool) bool {
	if len(s) == begin {
		return true
	}
	// terminate
	if res, exist := memo[begin]; exist {
		return res
	}
	// for choice in choiceList
	for word := range wordSet {
		if strings.HasPrefix(s[begin:], word) && dfs(s, begin+len(word), memo, wordSet) {
			memo[begin] = true
			return true
		}
	}
	memo[begin] = false
	return false
}
