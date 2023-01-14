package leetcode

import (
	"container/heap"
	"sort"
)

// https://leetcode.cn/problems/sliding-window-maximum/
// [1,3,-1,-3,5,3,6,7], k = 3
// [1,3,-1],-3,5,3,6,7
// 1,[3,-1,-3],5,3,6,7
// 1,3,[-1,-3,5],3,6,7
// 1,3,-1,[-3,5,3],6,7
// 1,3,-1,-3,[5,3,6],7
// 1,3,-1,-3,5,[3,6,7]

var a []int

// 大顶堆
type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}

func (h *hp) Push(x interface{}) {
	idx := x.(int)
	h.IntSlice = append(h.IntSlice, idx)
}

func (h *hp) Pop() interface{} {
	length := len(h.IntSlice)
	val := h.IntSlice[length-1]
	h.IntSlice = h.IntSlice[:length-1]
	return a[val]
}

func maxSlidingWindow(nums []int, k int) []int {
	// param check
	if len(nums) == 0 || k == 1 {
		return nums
	}

	var (
		pq     = &hp{}
		length = len(nums)
		result = make([]int, 0, length-k+1)
	)
	a = nums

	heap.Init(pq)
	// process
	for i := 0; i < k; i++ {
		heap.Push(pq, i)
	}
	result = append(result, a[pq.IntSlice[0]])

	for i := k; i < length; i++ {
		heap.Push(pq, i)
		for pq.IntSlice[0] < i-k+1 {
			heap.Pop(pq)
		}
		// peek = pq.IntSlice[0]
		result = append(result, a[pq.IntSlice[0]])
	}
	return result
}
