package leetcode_2009

import "sort"

// https://leetcode.cn/problems/minimum-number-of-operations-to-make-array-continuous/
func minOperations(nums []int) int {
	// 排序 + 去重 + 滑动窗口
	// 1. 枚举左端点, 窗口长度未 n. 则区间为: [left, left + n - 1]
	// 2. 去重, 重复数字 是一定要有一次操作的.  额外的操作次数就是  n - len(set)

	// 数字去重
	n := len(nums)
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	result := n
	arr := sortNums(set)
	j := 0
	for i, left := range arr {
		// 枚举左端点num
		right := left + n - 1
		for j < len(arr) && arr[j] <= right {
			j++
		}
		result = min(result, n-(j-i))
	}
	return result
}

func sortNums(set map[int]bool) []int {
	arr := make([]int, 0, len(set))
	for item := range set {
		arr = append(arr, item)
	}
	sort.Ints(arr)
	return arr
}
