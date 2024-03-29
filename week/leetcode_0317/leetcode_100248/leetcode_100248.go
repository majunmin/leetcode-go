package leetcode_100248

// https://leetcode.cn/contest/weekly-contest-389/problems/existence-of-a-substring-in-a-string-and-its-reverse/
func isSubstringPresent(s string) bool {
	if len(s) < 2 {
		return false
	}
	// 判断字符串中是否存在回文子串
	subStrSet := make(map[string]bool)
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return true
		}
		sub := s[i-1 : i+1]
		if subStrSet[sub] {
			return true
		}
		subStrSet[string([]byte{s[i], s[i-1]})] = true
	}
	return false
}
