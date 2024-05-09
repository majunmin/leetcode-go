package leetcode_2673

// https://leetcode.cn/problems/make-costs-of-paths-equal-in-a-binary-tree/
// 自底向上 + 贪心的解法  https://leetcode.cn/problems/make-costs-of-paths-equal-in-a-binary-tree/solutions/2656293/shi-er-cha-shu-suo-you-lu-jing-zhi-xiang-65hk/
func minIncrements(n int, cost []int) int {
	// param check
	if n < 2 {
		return 0
	}
	var result int
	// i 的初始值 设置为 n-1 还是设置为 n -2 呢?  - 都可以.
	for i := n - 1; i > 0; i -= 2 {
		incr := abs(cost[i] - cost[i-1])
		result += incr
		cost[i/2] += incr
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
