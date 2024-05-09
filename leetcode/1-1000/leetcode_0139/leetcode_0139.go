package leetcode_0139

import "strings"

// https://leetcode.cn/problems/word-break/description/?envType=study-plan-v2&envId=top-100-liked
func wordBreak(s string, wordDict []string) bool {
	if len(s) < 1 {
		return true
	}
	return dpSolution(s, wordDict)
}

func dpSolution(s string, dict []string) bool {
	dp := make([]bool, len(s)+1)
	wordMap := make(map[string]bool)
	for _, word := range dict {
		wordMap[word] = true
	}

	dp[0] = true
	//
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			suffix := s[j:i]
			if wordMap[suffix] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}

/*
带记忆化的 dfs
*/
func dfsSolutionMem(s string, begin int, dict []string, idxSet map[int]struct{}) bool {
	for _, word := range dict {
		_, exist := idxSet[begin]
		if !exist && strings.HasPrefix(s, word) {
			nextIdx := begin + len(word)
			idxSet[nextIdx] = struct{}{}
			if dfsSolutionMem(s, nextIdx, dict, idxSet) {
				return true
			}
		}
	}
	return false
}

func dfsSolution(s string, wordDict []string) bool {
	// terminate
	if len(s) == 0 {
		return true
	}
	//
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) && wordBreak(word[len(word):], wordDict) {
			return true
		}
	}
	return false
}
