package leetcode_0169

import "sort"

// https://leetcode.cn/problems/majority-element/?envType=study-plan-v2&envId=top-100-liked
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)>>1]
}

// Boyer-Moore 投票算法
func solution1(nums []int) int {
	var cnt, candidate int
	for _, num := range nums {
		if cnt == 0 {
			candidate = num
		}
		if num == candidate {
			cnt += 1
		} else {
			cnt -= 1
		}
	}
	return candidate
}
