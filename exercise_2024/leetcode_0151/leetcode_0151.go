package leetcode_0151

import "strings"

// https://leetcode.cn/problems/reverse-words-in-a-string/?envType=study-plan-v2&envId=top-interview-150
func reverseWords(s string) string {
	var (
		result   []string
		sBuilder strings.Builder
	)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if sBuilder.Len() > 0 {
				result = append([]string{sBuilder.String()}, result...)
			}
			sBuilder.Reset()
		} else {
			sBuilder.WriteByte(s[i])

		}

	}
	if sBuilder.Len() > 0 {
		result = append([]string{sBuilder.String()}, result...)
		sBuilder.Reset()
	}

	if len(result) == 0 {
		return ""
	}
	return strings.Join(result, " ")
}
