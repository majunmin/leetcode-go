package leetcode_2527

// https://leetcode.cn/problems/find-xor-beauty-of-array/
// 题解: https://leetcode.cn/problems/find-xor-beauty-of-array/solutions/2050337/no6289-cha-xun-shu-zu-xor-mei-li-zhi-by-d5ylk/
func xorBeauty(nums []int) int {
	var res int
	for _, num := range nums {
		res ^= num
	}
	return res
}

func baoli(nums []int) int {
	var res int
	for _, n1 := range nums {
		for _, n2 := range nums {
			for _, n3 := range nums {
				res ^= (n1 | n2) & n3
			}
		}
	}
	return res
}
