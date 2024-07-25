package leetcode_2766

import "slices"

// https://leetcode.cn/problems/relocate-marbles/
func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	var (
		stones = make(map[int]bool)
		result = make([]int, 0, len(nums))
	)
	for _, num := range nums {
		stones[num] = true
	}
	for i := range moveFrom {
		from, to := moveFrom[i], moveTo[i]
		delete(stones, from)
		stones[to] = true
	}
	for num := range stones {
		result = append(result, num)
	}
	slices.Sort(result)
	return result
}
