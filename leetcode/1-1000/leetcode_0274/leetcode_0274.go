package leetcode_0274

import "sort"

// https://leetcode.cn/problems/h-index/
func hIndex(citations []int) int {
	sort.Ints(citations)
	h, i := 0, len(citations)-1
	for i >= 0 && citations[i] > h {
		h++
		i--
	}
	return h
}

// 二分搜索
func hIndex2(citations []int) int {
	var (
		n           = len(citations)
		left, right = 0, n
	)
	// find  符合条件的 最大值
	for left < right {
		mid := left + (right-left+1)>>1
		if check(citations, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func check(citations []int, mid int) bool {
	var cnt int
	for _, num := range citations {
		if num >= mid {
			cnt++
		}
	}
	return cnt >= mid
}
