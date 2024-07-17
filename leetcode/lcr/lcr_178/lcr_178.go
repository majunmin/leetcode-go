package lcr_178

// https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/description/
func trainingPlan(actions []int) int {
	var result int
	for i := 0; i < 32; i++ {
		var cnt int
		for _, num := range actions {
			// cnt += num >> i & 1
			if num>>i&1 == 1 {
				cnt++
			}
		}
		if cnt%3 > 0 {
			result |= 1 << i
		}
	}
	return result
}
