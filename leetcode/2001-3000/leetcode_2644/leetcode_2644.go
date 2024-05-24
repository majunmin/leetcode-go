package leetcode_2644

import (
	"math"
	"slices"
)

// https://leetcode.cn/problems/find-the-maximum-divisibility-score/?envType=daily-question&envId=2024-05-18
func maxDivScore(nums []int, divisors []int) int {
	var (
		maxScore   int
		divisor    = math.MaxInt
		divisorSet = make(map[int]bool)
	)

	for _, item := range divisors {
		if divisorSet[item] {
			continue
		}
		divisorSet[item] = true
		// cal score
		var score int
		for _, num := range nums {
			if num%item == 0 {
				score++
			}
		}
		if score > maxScore {
			maxScore = score
			divisor = item
		} else if score == maxScore {
			divisor = min(divisor, item)
		}
	}
	return divisor
}

func maxDivScore2(nums []int, divisors []int) int {
	var (
		mx         = slices.Max(nums)
		maxScore   int
		divisor    = math.MaxInt
		divisorSet = make(map[int]bool)
		numCnt     = make([]int, mx)
	)
	for _, num := range nums {
		numCnt[num]++
	}

	for _, item := range divisors {
		if divisorSet[item] {
			continue
		}
		divisorSet[item] = true
		var score int
		for j := item; j <= mx; j += item {
			score += nums[j]
		}
		if score > maxScore {
			divisor = item
		} else if score == maxScore {
			divisor = min(divisor, item)
		}
	}
	return divisor
}
