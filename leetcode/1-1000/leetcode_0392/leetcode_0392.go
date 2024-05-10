package leetcode_0392

// https://leetcode.cn/problems/is-subsequence/?envType=study-plan-v2&envId=top-interview-150
func isSubsequence(s string, t string) bool {
	// param check
	n1, n2 := len(s), len(t)
	if n1 > n2 {
		return false
	}
	p1, p2 := 0, 0
	for p1 < n1 && p2 < n2 {
		// 字符匹配, p1 ++
		if s[p1] == t[p2] {
			p1++
		}
		p2++
	}
	return p1 == n1-1
}
