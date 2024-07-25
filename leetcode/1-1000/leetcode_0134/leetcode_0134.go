package leetcode_0134

import "math"

// https://leetcode.cn/problems/gas-station/?envType=study-plan-v2&envId=top-interview-150
func canCompleteCircuit(gas []int, cost []int) int {
	var (
		n        = len(gas)
		minSpare = math.MaxInt
		minIdx   int
		spare    int
	)
	for i := 0; i < n; i++ {
		spare += gas[i] - cost[i]
		if spare < minSpare {
			minSpare = spare
			minIdx = i
		}
	}
	return minIdx
}
