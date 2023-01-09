package offer_059

import (
	"container/heap"
	"sort"
)

// https://leetcode.cn/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/?envType=study-plan&id=lcof
func maxSlidingWindow(nums []int, k int) []int {
	return decreaseQueueSolution(nums, k)
}

// 单调递减队列
func decreaseQueueSolution(nums []int, k int) []int {
	n := len(nums)
	if k <= 0 || n < k {
		return nil // panic
	}
	if k == 1 {
		return nums
	}
	// 存储 index
	queue := make([]int, 0, 3)
	for i := 0; i < k; i++ {
		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	result := make([]int, n-k+1)
	result[0] = nums[queue[0]] // head
	for i := k; i < n; i++ {
		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		// 每次讲窗口左边的 idx 排除
		if queue[0] < i-k+1 {
			queue = queue[1:]
		}

		result[i-k+1] = nums[queue[0]]
	}
	return result
}

// heapSolution 使用  对大根堆  解决
func heapSolution(nums []int, k int) []int {
	// param check
	n := len(nums)
	if k <= 0 || n < k {
		return nil // panic
	}
	a = nums
	result := make([]int, n-k+1)
	h := &hp{make([]int, 0)}
	//heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, i)
	}
	result[0] = nums[h.IntSlice[0]]

	// 遍历后面的
	for i := k; i < n; i++ {
		heap.Push(h, i)
		for h.IntSlice[0] <= i-k {
			heap.Pop(h)
		}
		result[i-k+1] = nums[h.IntSlice[0]]
	}
	return result
}

var (
	a []int
)

type hp struct {
	sort.IntSlice
}

func (h *hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}
