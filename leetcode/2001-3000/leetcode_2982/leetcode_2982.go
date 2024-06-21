package leetcode_2982

import (
	"slices"
)

// https://leetcode.cn/problems/find-longest-special-substring-that-occurs-thrice-ii/?envType=daily-question&envId=2024-05-30
func maximumLength(s string) int {
	// 分组讨论
	var (
		n      = len(s)
		cnts   = make([][]int, 26)
		cnt    int
		result int
	)
	// result = max(a[0]-2, min(a[0]-1, a[1]), a[2])

	for i := 0; i < n; i++ {
		cnt++
		if i == n-1 || s[i] != s[i+1] {
			cnts[s[i]-'a'] = append(cnts[s[i]-'a'], cnt)
			cnt = 0
		}
	}
	for i := 0; i < 26; i++ {
		slices.SortFunc(cnts[i], func(a, b int) int {
			return b - a
		})
		if len(cnts) > 0 {
			result = max(result, cnts[i][0]-2)
		}
		if len(cnts) > 1 {
			result = max(result, min(cnts[i][0]-1, cnts[i][1]))
		}
		if len(cnts) > 2 {
			result = max(result, cnts[i][2])
		}
	}
	return result
}
