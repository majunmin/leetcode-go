package leetcode_0260

// https://leetcode.cn/problems/single-number-iii/
func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	// a ^b ^a = b
	lsb := xorSum & (xorSum - 1) // 取 xorsum 最后一位1
	// 分组计算
	var a, b int
	for _, num := range nums {
		if lsb&num == lsb {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
