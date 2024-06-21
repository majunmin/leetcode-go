package leetcode_0215

import (
	"container/heap"
	"sort"
)

// hp_solution: 建立容量为k的大顶堆, 将数组元素入堆, 然后取出 pop 堆k次，找到第k大值.
// solution2: 快速排序

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

// solution2
// 超时啦...
func findKthLargest2(nums []int, k int) int {
	// param check
	if k < 0 || k > len(nums) {
		return -1
	}
	pivot := nums[len(nums)-1]
	var cnt int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= pivot {
			nums[i], nums[cnt] = nums[cnt], nums[i]
			cnt++
		}
	}
	nums[cnt], nums[len(nums)-1] = nums[len(nums)-1], nums[cnt]
	if k == cnt+1 {
		return nums[cnt] // pivot
	}
	if k < cnt+1 {
		return findKthLargest2(nums[:cnt], k)
	}
	// k > cnt + 1
	return findKthLargest2(nums[cnt+1:], k-cnt-1)
}
