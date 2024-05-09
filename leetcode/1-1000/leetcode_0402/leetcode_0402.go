package leetcode_0402

import "strings"

// https://leetcode.cn/problems/remove-k-digits/
func removeKdigits(num string, k int) string {
	// param check
	if k == 0 {
		return num
	}

	stack := make([]byte, 0)
	remain := len(num) - k
	for i := range num {
		ch := num[i]
		for k > 0 && len(stack) > 0 && ch <= stack[len(stack)-1] {
			// pop
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, ch)
	}

	res := strings.TrimLeft(string(stack[:remain]), "0")
	if len(res) == 0 {
		return "0"
	}
	return res
}
