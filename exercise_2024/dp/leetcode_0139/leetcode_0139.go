package leetcode_0139

import (
	"bytes"
)

// https://leetcode.cn/problems/word-break/?envType=study-plan-v2&envId=top-100-liked
func wordBreak(s string, wordDict []string) bool {

	// param check
	if len(s) == 0 {
		return true
	}
	dp := make([]bool, len(s)+1)
	// dp[0] = true
	dp[0] = true
	sBytes := []byte(s)
	for i := 0; i < len(dp); i++ {
		if dp[i] == false {
			continue
		}
		for j := 0; j < len(wordDict); j++ {
			word := wordDict[j]
			if i+len(word) > len(s) {
				continue
			}
			if bytes.Equal(sBytes[i:i+len(word)], []byte(word)) {
				dp[i+len(word)] = true
			}
		}
	}
	return dp[len(dp)-1]
}
