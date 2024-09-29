package leetcode_0008

// https://leetcode.cn/problems/string-to-integer-atoi/description/
func myAtoi(s string) int {
	// 处理前导 空格
	var (
		res  int
		sign bool // 正数 为 false
		idx  int
	)
	for ; idx < len(s); idx++ {
		if s[idx] != ' ' {
			break
		}
	}
	// 处理符号位
	if s[idx] == '+' || s[idx] == '-' {
		sign = s[idx] == '-'
		idx++
	}
	//处理数字
	for ; idx < len(s); idx++ {
		if !isDigit(s[idx]) {
			break
		}
		res += res*10 + int(s[idx]-'0')
	}
	if sign {
		return -res
	}
	return res
}

func isDigit(x byte) bool {
	return x >= '0' && x <= '9'
}
