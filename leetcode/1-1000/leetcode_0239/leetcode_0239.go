package leetcode_0239

import (
	"container/heap"
	"sort"
)

var a []int

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	peek := len(h.IntSlice) - 1
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return a[peek]
}

// 滑动窗口最大值
//https://leetcode-cn.com/problems/sliding-window-maximum/
func maxSlidingWindow(nums []int, k int) []int {
	// paramCheck

	a = nums
	h := &hp{make([]int, 0, k)}
	for i := 0; i < k-1; i++ {
		h.IntSlice = append(h.IntSlice, i)
	}
	heap.Init(h)

	result := make([]int, 0, len(nums)-k+1)
	for i := k - 1; i < len(nums); i++ {
		heap.Push(h, i)
		for h.IntSlice[0] < i-k+1 {
			heap.Pop(h)
		}
		result = append(result, nums[h.IntSlice[0]])
	}
	return result
}

// 双端队列解法
func queueSolution(nums []int, k int) []int {
	// param check
	if k <= 0 {
		return nums
	}
	if k == 1 {
		return nums
	}

	// 双端队列
	queue := make([]int, 0, k)
	for i := 0; i < k-1; i++ {
		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	res := make([]int, 0, len(nums)-k+1)
	for i := k - 1; i < len(nums); i++ {
		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if queue[0] < i-k+1 {
			queue = queue[1:]
		}

		res = append(res, nums[queue[0]])
	}
	return res
}
