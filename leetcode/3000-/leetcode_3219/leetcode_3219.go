package leetcode_3219

import (
	"slices"
)

// https://leetcode.cn/problems/minimum-cost-for-cutting-cake-ii/?envType=daily-question&envId=2024-12-31
// 贪心解法
func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int64 {
	slices.Sort(horizontalCut)
	slices.Sort(verticalCut)
	var ans int64
	var i, j int
	for _ = range m + n - 2 {
		if j == n-1 || i < m-1 && horizontalCut[i] < verticalCut[j] {
			ans += int64(horizontalCut[i] * (n - j)) // 上下连边
			i++
		} else {
			ans += int64(verticalCut[j] * (m - i)) // 左右连边
			j++
		}
	}
	return ans
}

// 最小生成树(合并过程)
func minimumCostSolution2(m int, n int, horizontalCut []int, verticalCut []int) int64 {
	slices.Sort(horizontalCut)
	slices.Sort(verticalCut)
	var ans int
	var i, j int
	// 总共的合并次数: (m-1) + (n-1) = m+n-2
	for _ = range m + n - 2 {
		if horizontalCut[i] <= verticalCut[j] {
			ans += horizontalCut[i] * (n - j)
			i++
		} else {
			ans += verticalCut[j] * (m - i)
			j++
		}
	}
	return int64(ans)
}
