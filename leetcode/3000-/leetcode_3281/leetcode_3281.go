package leetcode_3281

import (
	"sort"
)

// https://leetcode.cn/problems/maximize-score-of-numbers-in-ranges/
func maxPossibleScore(start []int, d int) int {
	// 所选数字的 最小值 和 最大值 可以确定.
	// 枚举得分, 用二分法.
	// ...2,3,4,5,6,7,8,9....x,x+1,x+2...
	// 如果最大得分是 score, 那么 <=score 的得分都是满足条件的, > score 的得分都是不满足条件的. 满足单调性.
	// 需要有个函数判断得分score 是否满足条件: check(start, d, score)
	if len(start) < 2 {
		return 0
	}
	sort.Ints(start)
	var (
		// 区间最小值: start[0], 区间最大值:start[len(start)-1] + d
		maxScore = start[len(start)-1] + d - start[0]
		l, r     = -1, maxScore + 1
	)
	for l+1 < r {
		mid := l + (r-l)>>1
		if check(start, d, mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

func check(start []int, d int, score int) bool {
	prevScore := start[0]
	for _, s := range start[1:] {
		// ----[s--s+d]----
		nextScore := prevScore + score
		if nextScore <= s {
			prevScore = s
			continue
		}
		if nextScore <= s+d {
			prevScore = nextScore
			continue
		}
		return false
	}
	return true
}
