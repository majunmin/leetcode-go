package leetcode_0012

import "strings"

var (
	valueSymbols = []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
)

// https://leetcode.cn/problems/integer-to-roman/?envType=study-plan-v2&envId=top-interview-150
func intToRoman(num int) string {
	var sb strings.Builder
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			sb.WriteString(vs.symbol)
			continue
		}
		if num == 0 {
			break
		}
	}
	return sb.String()
}
