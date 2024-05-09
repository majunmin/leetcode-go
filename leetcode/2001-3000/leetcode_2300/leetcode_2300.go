package leetcode_2300

import "sort"

// https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/
func successfulPairs(spells []int, potions []int, success int64) []int {
	//
	var (
		n      = len(spells)
		result = make([]int, n)
	)
	sort.Ints(potions)
	for i := 0; i < n; i++ {
		idx := binarySearch(potions, spells[i], success)
		result[i] = idx
	}
	return result
}

func binarySearch(potions []int, spell int, success int64) int {
	var (
		left, right = 0, len(potions) - 1
	)
	for left < right {
		mid := left + (right-left)>>1
		if potions[mid]*spell < int(success) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return len(potions) - right
}
