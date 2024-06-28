package _024_06_02

import "slices"

// https://leetcode.cn/problems/lexicographically-minimum-string-after-removing-stars/
func clearStars(s string) string {
	var (
		stacks = make([][]int, 26)
	)
	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			stacks[s[i]-'a'] = append(stacks[s[i]-'a'], i)
			continue
		}
		// pop一个字符
		for j, ss := range stacks {
			if len(ss) > 0 {
				stacks[j] = ss[:len(ss)-1]
				break
			}
		}
	}
	// 还原字符
	idxs := make([]int, 0)
	for _, ss := range stacks {
		idxs = append(idxs, ss...)
	}
	slices.SortFunc(idxs, func(a, b int) int {
		return a - b
	})
	result := make([]byte, 0, len(idxs))
	for _, idx := range idxs {
		result = append(result, s[idx])
	}
	return string(result)
}
