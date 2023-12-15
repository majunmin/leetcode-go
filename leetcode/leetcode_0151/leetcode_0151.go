package leetcode_0151

import (
	"slices"
	"strings"
)

// https://leetcode.cn/problems/reverse-words-in-a-string/
func reverseWords(s string) string {
	wordList := make([]string, 0, 16)

	wordLen := 0
	var sBuilder strings.Builder
	for i := range s {
		if s[i] == ' ' && wordLen > 0 {
			wordList = append(wordList, sBuilder.String())
			sBuilder.Reset()
			wordLen = 0
		}
		if s[i] == ' ' {
			continue
		}
		wordLen++
		sBuilder.WriteByte(s[i])
	}
	if wordLen > 0 {
		wordList = append(wordList, sBuilder.String())
		sBuilder.Reset()
	}

	slices.Reverse(wordList)
	return strings.Join(wordList, " ")
}
