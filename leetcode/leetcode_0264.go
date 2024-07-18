package leetcode

import (
	"strings"
)

func convertToBase7(num int) string {
	if num < 0 {
		return "-" + convertToBase7(num)
	}
	var sb strings.Builder
	for num > 0 {
		cur := num % 7
		num /= 7
		sb.WriteByte(byte(cur + '0'))
	}
	return sb.String()
}

func reverseString(s string) string {
	raw := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		raw[l], raw[r] = raw[r], raw[l]
		l++
		r--
	}
	return string(raw)
}
