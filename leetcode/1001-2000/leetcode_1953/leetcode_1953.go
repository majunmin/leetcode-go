package leetcode_1953

import "sort"

// https://leetcode.cn/problems/maximum-number-of-weeks-for-which-you-can-work/
// 贪心法.
// 贪心法 最难的点是如何证明?
func numberOfWeeks(milestones []int) int64 {
	//param check
	if len(milestones) == 0 {
		return 0
	}
	var (
		n       = len(milestones)
		longest int
		rest    int
	)

	sort.Ints(milestones)
	longest = milestones[n-1]
	for i := 0; i < n-1; i++ {
		rest += milestones[i]
	}
	if longest <= rest+1 {
		return int64(longest + rest)
	}
	return int64(rest*2 + 1)
}
