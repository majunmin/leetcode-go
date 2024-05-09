package leetcode_0191

//https://leetcode.cn/problems/number-of-1-bits/
func hammingWeight(num uint32) int {
	n := num
	var cnt int
	for n > 0 {
		if n&1 == 1 {
			cnt += 1
		}
		n >>= 1
	}
	return cnt
}
