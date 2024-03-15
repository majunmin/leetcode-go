package leetcode_2129

import "strings"

// https://leetcode.cn/problems/capitalize-the-title/
func capitalizeTitle(title string) string {
	var sb strings.Builder
	var word []byte
	for i := range title {
		if title[i] == ' ' {
			// terminate
			if len(word) > 2 {
				word[0] -= 32
			}
			sb.Write(word)
			sb.WriteByte(' ')
			// clear word
			word = word[:0]
			continue
		}
		item := title[i]
		if item >= 'A' && item <= 'Z' {
			// 'a' - 'A' = 32
			item += 32
		}
		word = append(word, byte(item))

	}
	if len(word) > 0 {
		if len(word) > 2 {
			word[0] -= 32
		}
		sb.Write(word)
	}
	return sb.String()
}
