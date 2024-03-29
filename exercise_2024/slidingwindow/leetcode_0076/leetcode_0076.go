package leetcode_0076

import "math"

// https://labuladong.online/algo/essential-technique/sliding-window-framework/
// https://leetcode.cn/problems/minimum-window-substring/description/
func minWindow(s string, t string) string {

	needs := make(map[byte]int)
	for i := range t {
		needs[t[i]]++
	}
	var valid, l, r int
	length := math.MaxInt
	window := make(map[byte]int)
	var result string
	for r < len(s) {
		//扩大窗口
		window[s[r]]++
		if window[s[r]] == needs[s[r]] {
			valid++
		}
		r++

		//shrink window
		for valid == len(needs) {
			// 存在包含 t 的子串
			if r-l < length {
				result = s[l:r]
				length = r - l
			}
			//shrink window
			d := s[l]
			l++
			window[d]--
			if _, exist := needs[d]; exist {
				if window[d]+1 == needs[d] {
					valid--
				}
			}
		}
	}
	return result
}
