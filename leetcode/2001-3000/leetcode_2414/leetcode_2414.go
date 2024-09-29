package leetcode_2414

// https://leetcode.cn/problems/length-of-the-longest-alphabetical-continuous-substring/
func longestContinuousSubstring(s string) int {
	// 滑动窗口解法
	if len(s) < 2 {
		return len(s)
	}
	var (
		r      = 0
		l      int
		n      = len(s)
		result = 1
	)
	for ; r < n; r++ {
		if s[r]-s[r-1] == 1 {
			result = max(result, r-l+1)
		} else {
			l = r
		}
	}
	return result
}
