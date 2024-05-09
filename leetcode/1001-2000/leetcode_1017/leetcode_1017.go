package leetcode_1017

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/convert-to-base-2/
func baseNeg2(n int) string {
	if n == 1 || n == 0 {
		return strconv.Itoa(n)
	}
	var res strings.Builder
	for n != 0 {
		if n%-2 == 0 {
			// n 为偶数
			res.WriteByte('0')
		} else {
			n -= 1 // 将最低位的值抹去后,使得n变为偶数, 在进行除负二操作时是整除
			// n 为奇数, 当前n的最低位f为1 (tips: 虽然 n % -2 可能为1或-1，但不影响,都是最低位f=1)
			res.WriteByte('1')
		}
		//将n的-2进制 右移一位
		n /= -2
	}
	return reverse(res.String())
}

func reverse(s string) string {
	rawBytes := []byte(s)
	l, r := 0, len(rawBytes)-1
	for l < r {
		rawBytes[l], rawBytes[r] = rawBytes[r], rawBytes[l]
		l++
		r--
	}
	return string(rawBytes)
}

// 十进制转 n 进制
func baseX(n, base int) string {
	if n == 0 || n == 1 {
		return strconv.Itoa(n)
	}
	var sb strings.Builder
	for n != 0 {
		remin := n % base
		if remin == 0 {
			sb.WriteByte('0')
		} else {
			if remin < 0 {
				// 说明x一定 < 0
				//将remin 转化为正数范围内的数字
				remin -= base // 将最后一位抹零，保证后续是整除，从而保证不同语言通用
			}
			n -= remin

			// remain >= 0
			if remin > 10 {
				sb.WriteByte(byte('A' + remin - 10))
			} else {
				sb.WriteByte(byte('0' + remin))
			}
		}
		n /= base
	}
	return reverse(sb.String())
}
