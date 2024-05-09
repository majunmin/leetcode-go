package leetcode_1318

// https://leetcode.cn/problems/minimum-flips-to-make-a-or-b-equal-to-c/
func minFlips(a int, b int, c int) int {
	var result int
	for i := 0; i < 32; i++ {
		abit := (a >> i) & 1
		bbit := (b >> i) & 1
		cbit := (c >> i) & 1
		if cbit == 1 {
			result += 1 - (abit | bbit)
		} else {
			result += abit + bbit
		}
	}
	return result
}
