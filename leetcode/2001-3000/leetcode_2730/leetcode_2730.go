package leetcode_2730

// https://leetcode.cn/problems/find-the-longest-semi-repetitive-substring/
func longestSemiRepetitiveSubstring(s string) int {
	// param check
	if len(s) < 3 {
		return len(s)
	}

	var (
		cnt    int //窗口内连续字符
		maxLen int
		prevl  int
	)

	l, r := 0, 1
	for r < len(s) {
		if s[r] == s[r-1] {
			cnt++
		}
		maxLen = max(maxLen, r-l)
		// if cnt < 2
		if cnt < 2 {
			r++
			continue
		}
		for l < r {
			if l > prevl && s[l] == s[l-1] {
				cnt--
				prevl = l
				break
			}
			l++
		}
		r++
	}
	return maxLen
}
