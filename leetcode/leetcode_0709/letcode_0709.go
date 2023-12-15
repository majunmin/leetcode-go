package leetcode_0709

import "strings"

func toLowerCase(s string) string {
	if len(s) == 0 {
		return ""
	}
	// 'a' - 'A' = 32
	var sBuilder strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			sBuilder.WriteByte(s[i] + 'a' - 'A')
			continue
		}
		sBuilder.WriteByte(s[i])
	}
	return sBuilder.String()
}
