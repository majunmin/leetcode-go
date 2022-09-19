package offer_005

import (
	"strings"
)

// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/?envType=study-plan&id=lcof
func replaceSpace(s string) string {
	//return strings.ReplaceAll(s, " ", "%20")
	var sb strings.Builder
	for _, r := range []rune(s) {
		if r == ' ' {
			sb.WriteString("%20")
		}
		sb.WriteRune(r)
	}
	return sb.String()
}
