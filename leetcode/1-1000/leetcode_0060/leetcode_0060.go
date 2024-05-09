package leetcode_0060

import "strings"

// https://leetcode.cn/problems/permutation-sequence/
func getPermutation(n int, k int) string {
	// 注意: 相当于在 n 个数字的全排列中找到下标为 k - 1 的那个数, 因此 k 先减 1
	k--
	var (
		sb    strings.Builder
		queue = make([]int, 0, n)
		fib   = make([]int, n)
	)
	fib[0] = 1
	for i := 1; i < n; i++ {
		fib[i] = fib[i-1] * i
	}
	for i := 1; i <= n; i++ {
		queue = append(queue, i)
	}
	for i := n - 1; i >= 0; i-- {
		idx := k / fib[i]
		sb.WriteByte(byte(queue[idx]) + '0')
		queue = append(queue[:idx], queue[idx+1:]...)
		k -= idx * fib[i]
	}
	return sb.String()
}
