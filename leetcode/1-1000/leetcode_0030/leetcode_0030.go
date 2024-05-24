package leetcode_0030

import (
	"sort"
	"strings"
)

// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/?envType=study-plan-v2&envId=top-interview-150
func findSubstring(s string, words []string) []int {
	//param check
	if len(words) == 0 || len(s) == 0 {
		return nil
	}
	// 固定窗口问题
	sort.Strings(words)
	var (
		target     = strings.Join(words, "")
		itemSize   = len(words[0])
		windowSize = len(target)
		result     = make([]int, 0, len(s))
	)
	if len(s) < windowSize {
		return nil
	}
	left, right := 0, windowSize-1 // 左闭右闭区间
	for left <= len(s)-windowSize {
		pattern := parseWord(s[left:right+1], itemSize)
		if pattern == target {
			result = append(result, left)
		}
	}
	return result
}

// 将s 按照 size 长度进行切片
func parseWord(s string, size int) string {
	items := make([]string, 0)
	for i := 0; i < len(s); i += size {
		items = append(items, s[i:i+size])
	}
	sort.Strings(items)
	return strings.Join(items, "")
}
