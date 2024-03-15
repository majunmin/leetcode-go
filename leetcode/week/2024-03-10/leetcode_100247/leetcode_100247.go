package leetcode_100247

import (
	"slices"
)

// https://leetcode.cn/contest/weekly-contest-388/problems/maximize-happiness-of-selected-children/
func maximumHappinessSum(happiness []int, k int) int64 {
	if k == 0 {
		return 0
	}
	if k > len(happiness) {
		panic("happiness is not enough")
	}
	// 贪心， 每次只挑选 幸福值最大的孩子
	slices.SortFunc(happiness, func(a, b int) int {
		return b - a
	})
	//
	var (
		result int64 //最大幸福值
		decr   int   // 本次挑选要减少的幸福值
	)
	for i := 0; i < k; i++ {
		result += int64(max(happiness[i]-decr, 0))
	}
	return result
}
