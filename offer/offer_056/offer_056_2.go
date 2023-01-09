package offer_056

// https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/?envType=study-plan&id=lcof
func singleNumber(nums []int) int {
	// 统计位数
	count := make([]int, 32) // 统计 每位1 的个数 //
	for _, num := range nums {
		for j := 0; j < 32; j++ {
			if (num>>j)&1 == 1 { // 判断 num 第 j 位是否为1
				count[j] += 1
			}
		}
	}

	var result int
	for i := 31; i >= 0; i-- {
		result <<= 1
		if count[i]%3 == 1 {
			result += 1
		}
	}
	return result
}
