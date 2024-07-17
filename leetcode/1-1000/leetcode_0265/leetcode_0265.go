package leetcode_0265

import "math"

// https://leetcode.ca/all/265.html
func minCostII(costs [][]int) int {
	if len(costs) == 0 || len(costs[0]) == 0 {
		return 0
	}
	var (
		n  = len(costs[0])
		dp = make([]int, n)
	)
	for _, cost := range costs {
		for j := range cost {
			incr := math.MaxInt
			for k := range cost {
				if k == j {
					continue
				}
				incr = min(incr, cost[k])
			}
			dp[j] += incr
		}
	}
	return min(dp[n])
}
