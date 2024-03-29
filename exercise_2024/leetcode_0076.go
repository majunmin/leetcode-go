package exercise_2024

import "math"

// https://leetcode.cn/problems/minimum-window-substring/?envType=study-plan-v2&envId=top-100-liked
func minWindow(s string, t string) string {

	// param check
	//
	needCnt := make(map[byte]int)
	for i := range t {
		needCnt[t[i]]++
	}
	var (
		l, r   int
		validC int

		resultLen = math.MaxInt
		result    string

		window = make(map[byte]int, len(needCnt))
	)

	for r < len(s) {
		window[s[r]]++
		if needCnt[s[r]] == window[s[r]] {
			validC++
		}
		r++
		// 找到符合条件的 了
		if validC == len(needCnt) {
			// shrink window
			for l < r {
				if window[s[l]] == needCnt[s[l]] {
					validC--
					if r-l < resultLen {
						resultLen = r - l
						result = s[l:r]
					}
					window[s[l]]--
					l++
					break
				}
				window[s[l]]--
				l++
			}
		}
	}
	return result
}
