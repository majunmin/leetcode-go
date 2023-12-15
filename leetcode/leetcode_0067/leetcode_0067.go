package leetcode_0067

import (
	"strings"
)

func addBinary(a string, b string) string {
	// param check
	if len(a) == 0 || len(b) == 0 {
		panic("invalid param")
	}
	var buf strings.Builder
	byteOfa, byteOfb := []byte(a), []byte(b)
	la, lb := len(a), len(b)
	length := maxInt(la, lb)
	var carry int // 进位,初始值 0
	for i := 0; i < length; i++ {
		ia := 0
		if i < la {
			ia = int(byteOfa[la-1-i] - '0')
		}
		ib := 0
		if i < lb {
			ib = int(byteOfb[lb-1-i] - '0')
		}
		carry = ia + ib + carry
		buf.WriteByte(byte(carry%2) + '0')
		carry = carry / 2
	}
	//最后一位进位
	if carry == 1 {
		buf.WriteByte('1')
	}
	res := []byte(buf.String())
	reverse(res)
	return string(res)
}

func reverse(res []byte) {
	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
