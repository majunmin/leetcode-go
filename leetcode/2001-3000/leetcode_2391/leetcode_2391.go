package leetcode_2391

import "strings"

// 模拟
// https://leetcode.cn/problems/minimum-amount-of-time-to-collect-garbage/
func garbageCollection(garbage []string, travel []int) int {
	// param check
	if len(garbage) == 0 {
		return 0
	}
	var (
		garbageTypes = []string{"M", "P", "G"}
		result       int
	)

	// 分别收拾的travelCost  垃圾 G M P
	// 这里可以将 travel 预处理成前缀和 **
	for _, gType := range garbageTypes {
		var travelCost int
		var accumulate int
		for i := 1; i < len(garbage); i++ {
			// add travel cost
			accumulate += travel[i-1]
			if strings.Contains(garbage[i], gType) {
				travelCost = accumulate
			}

		}
		result += travelCost
	}
	// 收拾 garage cost
	for _, g := range garbage {
		result += len(g)
	}
	return result
}

func garbageCollection2(garbage []string, travel []int) int {
	// param check
	if len(garbage) == 0 {
		return 0
	}
	var (
		garbageTypes = []string{"M", "P", "G"}
		result       int
		travelPreSum = make([]int, len(travel)+1)
	)
	travelPreSum[0] = 0

	// 1. 收拾 garage cost
	for _, g := range garbage {
		result += len(g)
	}
	// 2. 预处理前缀和
	for i := 0; i < len(travel); i++ {
		travelPreSum[i+1] = travelPreSum[i] + travel[i]
	}

	// 分别收拾的travelCost  垃圾 G M P
	for _, gType := range garbageTypes {
		var travelCost int
		for i := 1; i < len(garbage); i++ {
			// 3. add travel cost
			if strings.Contains(garbage[i], gType) {
				travelCost = travelPreSum[i]
			}

		}
		result += travelCost
	}
	return result
}
