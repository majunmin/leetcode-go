package leetcode_0014

import "strings"

// https://leetcode.cn/problems/longest-common-prefix/description/
func longestCommonPrefix(strs []string) string {
	// param check
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = maxPrefix(prefix, strs[i])
	}
	return prefix
}

func maxPrefix(s1 string, s2 string) string {
	var i int
	var result strings.Builder
	for i < len(s1) && i < len(s2) {
		if s1[i] != s2[i] {
			break
		}
		result.WriteByte(s1[i])
		i++
	}
	return result.String()
}
