package lcr_017

import "math"

// https://leetcode.cn/problems/M1oyTv/
func minWindow(s string, t string) string {
	// 滑动窗口问题
	var (
		m, n = len(s), len(t)
	)
	if m < n {
		return ""
	}
	var (
		needs = make(map[byte]int)
		// t中 字符的种类数  len(needs)
		types  int
		result = ""
		resLen = math.MaxInt
		window = make(map[byte]int)
		l, r   int
	)

	for i := range t {
		needs[t[i]]++
	}
	for ; r < m; r++ {
		window[s[r]]++
		if window[s[r]] == needs[s[r]] {
			types++
		}

		for ; l <= r && types == len(needs); l++ {
			// shrink window
			if _, exist := needs[s[l]]; !exist {
				continue
			}
			window[s[l]]--
			if window[s[l]]+1 == needs[s[l]] {
				types--
				curLen := r - l + 1
				if curLen < resLen {
					resLen = curLen
					result = s[l : r+1]
				}
			}
		}
	}
	return result
}
