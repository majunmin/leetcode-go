package leetcode_0151

import "strings"

func reverseWords(s string) string {
	words := make([]string, 0)
	var preWord []byte
	for _, b := range s + " " {
		if b == ' ' {
			if len(preWord) > 0 {
				words = append(words, string(preWord))
			}
			preWord = preWord[:0]
			continue
		}
		preWord = append(preWord, byte(b))
	}
	reverse(words)
	return strings.Join(words, " ")
}

func reverse(words []string) {
	l, r := 0, len(words)-1
	for l < r {
		words[l], words[r] = words[r], words[l]
		l++
		r--
	}
}
