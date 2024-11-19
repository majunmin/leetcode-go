package leetcode_3226

func minChanges(n int, k int) int {
	// n: 11010
	// k: 01001
	// n^k: 10011
	// 1 of n should contains k
	if n == k {
		return 0
	}
	r1 := n ^ k
	if r1&k != 0 {
		return -1
	}
	return count1(r1)
}

func count1(n int) int {
	var res int
	for n > 0 {
		res++
		n = n & (n - 1)
	}
	return res
}
