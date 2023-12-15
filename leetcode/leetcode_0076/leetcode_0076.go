package leetcode_0076

// https://leetcode.cn/problems/minimum-window-substring/description/?envType=study-plan-v2&envId=top-100-liked
func minWindow(s string, t string) string {
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	needCnt := len(need)

	al, ar := 0, len(s)+1
	window := make(map[byte]int)
	l, r := 0, 0 // 滑动窗口
	for ; r < len(s); r++ {
		window[s[r]]++
		if window[s[r]] == need[s[r]] {
			needCnt--
		}
		for needCnt == 0 { // 此时包含了 全部的 字符
			//
			if r-l < ar-al {
				al, ar = l, r
			}
			l++
			// shrink window
			if window[s[l-1]] == need[s[l-1]] {
				needCnt++
				break
			}
		}
	}
	if ar-al == len(s)+1 {
		return ""
	}
	return s[al : ar+1]
}

// 滑动窗口通用解法1
func solution1(s string, t string) string {
	// param check
	need := make(map[byte]int)
	cnt := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	// process
	al, ar := 0, len(s)
	l, r := 0, 0
	for ; r < len(s); r++ {
		cnt[s[r]]++
		for check(cnt, need) {
			//  满足条件, shrink window
			if r-l < ar-al {
				al, ar = l, r
			}
			cnt[s[l]]--
			l++
		}
	}
	if ar-al < len(s) {
		return s[al:ar]
	}
	return ""
}

func check(window map[byte]int, need map[byte]int) bool {
	for b, cnt := range need {
		if window[b] < cnt {
			return false
		}
	}
	return true
}
