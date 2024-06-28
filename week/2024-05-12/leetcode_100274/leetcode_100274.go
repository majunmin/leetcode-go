package leetcode_100274

import "math"

func maximumEnergy(energy []int, k int) int {
	// param check
	result := math.MinInt
	for i := 0; i < len(energy); i++ {
		// if i < k {
		// 	result = max(result, energy[i])
		// 	continue
		// }
		if i >= k {
			energy[i] = max(energy[i], energy[i]+energy[i-k])
		}
		// 统计 result
		if i >= len(energy)-k {
			result = max(result, energy[i])
		}
	}
	return result
}
