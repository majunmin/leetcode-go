package leetcode_0881

import "sort"

// https://leetcode.cn/problems/boats-to-save-people/
// 1 <= people[i] <= limit
func numRescueBoats(people []int, limit int) int {
	// 贪心 + 双指针
	sort.Ints(people)
	l, r := 0, len(people)-1
	var result int
	for l < r {
		if people[l]+people[r] > limit {
			r--
		} else {
			r--
			l++
		}
		result++
	}
	// 加上最后剩下的一个人
	if l == r {
		result += 1
	}
	return result
}
