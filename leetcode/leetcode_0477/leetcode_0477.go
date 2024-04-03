package leetcode_0477

// https://leetcode.cn/problems/total-hamming-distance/
func totalHammingDistance(nums []int) int {
	count1 := make([]int, 32)
	var total int
	n := len(nums)
	for i := 0; i < 32; i++ {
		for _, num := range nums {
			num >>= i
			count1[i] += num & 1
		}
		total += count1[i] * (n - count1[i])
	}
	return total
}

// 简化简化在简化
func totalHammingDistance2(nums []int) int {
	var total int
	n := len(nums)
	for i := 0; i < 32; i++ {
		var cnt1 int
		for _, num := range nums {
			num >>= i
			cnt1 += num & 1
		}
		total += cnt1 * (n - cnt1)
	}
	return total
}
