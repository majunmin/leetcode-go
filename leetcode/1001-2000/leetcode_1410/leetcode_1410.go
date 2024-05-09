package leetcode_1410

import "strings"

// https://leetcode.cn/problems/html-entity-parser/
func entityParser(text string) string {

	flagMap := map[string]byte{
		"&amp;":   '&',
		"&nbsp;":  ' ',
		"&quot;":  '"',
		"&apos;":  '\'',
		"&gt;":    '>',
		"&lt;":    '<',
		"&frasl;": '/',
	}
	var result strings.Builder
	for i := 0; i < len(text); {
		if text[i] != '&' {
			result.WriteByte(text[i])
			i++
			continue
		}
		for k, v := range flagMap {
			if strings.HasPrefix(text[i:], k) {
				result.WriteByte(v)
				i += len(k)
				break
			}
		}
	}
	return result.String()
}
