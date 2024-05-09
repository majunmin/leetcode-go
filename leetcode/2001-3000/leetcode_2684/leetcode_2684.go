package leetcode_2684

// https://leetcode.cn/problems/maximum-odd-binary-number/
func maximumOddBinaryNumber(s string) string {
	// 01010101010
	rawBytes := []byte(s)
	// 先把所有的1 移到 最左边.， 最后将其其 变为奇数
	var l, r int
	// 移位,参考(移动0).
	for r < len(rawBytes) {
		if rawBytes[r] == '1' {
			rawBytes[r], rawBytes[l] = rawBytes[l], rawBytes[r]
			l++
		}
		r++
	}
	// 变为奇数
	rawBytes[l-1], rawBytes[len(rawBytes)-1] = rawBytes[len(rawBytes)-1], rawBytes[l-1]
	return string(rawBytes)
}
