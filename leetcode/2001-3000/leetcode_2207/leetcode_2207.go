package leetcode_2207

// https://leetcode.cn/problems/maximize-number-of-subsequences-in-a-string/
func maximumSubsequenceCount(text string, pattern string) int64 {
	var (
		result int // 一共有多少子序列
		p0Cnt  int
		p1Cnt  int
	)
	for i := range text {
		if text[i] == pattern[1] {
			result += p0Cnt
			p1Cnt++
		} else if text[i] == pattern[0] {
			p0Cnt++
		}
	}
	return int64(result + max(p0Cnt, p1Cnt))
}
