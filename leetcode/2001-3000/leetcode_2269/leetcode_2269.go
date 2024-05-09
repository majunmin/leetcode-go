package leetcode_2269

import "strconv"

// https://leetcode.cn/problems/find-the-k-beauty-of-a-number/
func divisorSubstrings(num int, k int) int {
	str := strconv.Itoa(num)
	var cnt int
	for i := 0; i <= len(str)-k; i++ {
		substr := str[i : i+k]
		subNum, _ := strconv.Atoi(substr)
		if subNum != 0 && num%subNum == 0 {
			cnt++
		}
	}
	return cnt
}
