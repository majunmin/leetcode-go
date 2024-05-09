package leetcode_0151

import (
	"slices"
	"strings"
)

// https://leetcode.cn/problems/reverse-words-in-a-string/
func reverseWords(s string) string {
	wordList := make([]string, 0, 16)
	var sBuilder strings.Builder
	s += " "
	for i := range s {
		if s[i] == ' ' {
			if sBuilder.Len() > 0 {
				wordList = append(wordList, sBuilder.String())
				sBuilder.Reset()
			}
			continue
		}
		sBuilder.WriteByte(s[i])
	}

	slices.Reverse(wordList)
	return strings.Join(wordList, " ")
}
