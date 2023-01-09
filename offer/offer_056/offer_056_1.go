package offer_056

// https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/?envType=study-plan&id=lcof
func singleNumbers(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nil
	}
	ret := 0
	// 如果 出现一次的 数字是 a b， 那么 全员异或的结果 就是  a ^ b
	for i := 0; i < length; i++ {
		ret ^= nums[i]
	}
	m := 1
	// 找到第一位 不相同的位置 , 目的是为了将数字分组
	for ret&m == 0 {
		m <<= 1
	}

	x, y := 0, 0 // 0^ n == n
	for i := range nums {
		if i&m == 0 {
			x ^= nums[i]
		} else {
			y ^= nums[i]
		}
	}
	return []int{x, y}
}
