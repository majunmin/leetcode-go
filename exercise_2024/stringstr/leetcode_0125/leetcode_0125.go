package leetcode_0125

import "strings"

// https://leetcode.cn/problems/valid-palindrome/
func isPalindrome(s string) bool {
	// param check
	if len(s) < 2 {
		return true
	}
	s = parseString(s)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	// left == right || left - right == 1
	return true
}

func parseString(s string) string {
	var sb strings.Builder
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' || (s[i] >= '0' && s[i] <= '9') {
			sb.WriteByte(s[i])
			continue
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			sb.WriteByte(s[i] - 'A' + 'a')
		}
	}
	return sb.String()
}
