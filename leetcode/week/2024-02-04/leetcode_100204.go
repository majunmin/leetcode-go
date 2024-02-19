package _024_02_04

import "strings"

// https://leetcode.cn/contest/weekly-contest-383/problems/minimum-time-to-revert-word-to-initial-state-i/
func minimumTimeToInitialState(word string, k int) int {
	// param check
	n := len(word)
	if k > n {
		// param invalid
		panic("invalid param, k > len(word)")
	}
	if k == n {
		return 1
	}
	//
	result := 1
	for i := 1; i*k < n; i++ {
		if strings.HasPrefix(word, word[i*k:]) {
			break
		}
		result++
	}
	return result
}
