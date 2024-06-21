package leetcode_0521

// https://leetcode.cn/problems/longest-uncommon-subsequence-i/
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}
