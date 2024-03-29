package _024_02_18

import "strings"

// https://leetcode.cn/contest/weekly-contest-385/problems/count-prefix-and-suffix-pairs-i/
func countPrefixSuffixPairs(words []string) int {
	var cnt int
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				cnt++
			}
		}
	}
	return cnt
}

func isPrefixAndSuffix(str1, str2 string) bool {
	return strings.HasPrefix(str2, str1) && strings.HasSuffix(str2, str1)
}
