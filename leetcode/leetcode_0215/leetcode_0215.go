package leetcode_0215

import (
	"container/heap"
	"sort"
)

type maxHeap struct {
	sort.IntSlice
}

func newMaxHeap() *maxHeap {
	return &maxHeap{}
}

func (h *maxHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *maxHeap) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *maxHeap) Pop() any {
	x := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return x
}

// https://leetcode.cn/problems/kth-largest-element-in-an-array/?envType=study-plan-v2&envId=top-100-liked
func findKthLargest(nums []int, k int) int {
	// param check
	if k < 1 || k > len(nums) {
		return -1
	}
	//
	mHp := newMaxHeap()
	for _, num := range nums {
		mHp.Push(num)
	}
	var kthNum int
	for i := 0; i < k; i++ {
		kthNum = heap.Pop(mHp).(int)
	}
	return kthNum
}
