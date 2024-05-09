package leetcode_2546

import "strings"

func makeStringsEqual(s string, target string) bool {
	// param checjk
	if len(s) != len(target) {
		return false
	}
	// 两个 字符串都包含1 || 两个字符串都不包含1 ,那么这两个字符串可以相互转化
	return strings.Contains(s, "1") && strings.Contains(target, "1") ||
		(!strings.Contains(s, "1") && !strings.Contains(target, "1"))
	//  return strings.Contains(s, "1") == strings.Contains(target, "1")
}
