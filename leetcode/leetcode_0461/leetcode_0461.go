package leetcode_0461

// https://leetcode.cn/problems/hamming-distance/
func hammingDistance(x int, y int) int {
	//return bits.OnesCount(uint(x ^ y))
	return countOne(x ^ y)
}

func countOne(y int) int {
	var result int
	for y > 0 {
		result++
		y &= y - 1
	}
	return result
}
