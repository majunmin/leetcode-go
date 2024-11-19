package leetcode_3258

// https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-i/?envType=daily-question&envId=2024-11-14
func countKConstraintSubstrings(s string, k int) int {
	var (
		n      = len(s)
		window = make([]int, 2)
		l, r   int
		result int
	)
	for r < n {
		window[s[r]-'0']++
		for l <= r && window[0] > k && window[1] > k {
			window[s[l]-'0']--
			l++
		}
		result += r - l + 1
	}
	return result
}
