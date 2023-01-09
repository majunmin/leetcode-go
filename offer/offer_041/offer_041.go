package offer_041

import (
	"container/heap"
	"sort"
)

// https://leetcode.cn/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/?envType=study-plan&id=lcof
type MedianFinder struct {
	maxHp, minHp hp //左边是 大顶堆 👈🏻 👉🏻  右边是 小顶堆
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	maxQ, minQ := &this.maxHp, &this.minHp
	if this.maxHp.Len() == 0 {
		heap.Push(maxQ, num)
		return
	}

	if this.maxHp.Len() == this.minHp.Len() {
		rightVal := heap.Pop(minQ).(int)
		heap.Push(minQ, rightVal)
		if num <= rightVal {
			heap.Push(maxQ, -num)
			return
		}
		heap.Push(maxQ, -heap.Pop(minQ).(int))
		heap.Push(minQ, num)

	} else {
		leftVal := -heap.Pop(maxQ).(int)
		heap.Push(maxQ, leftVal)
		if num >= leftVal {
			heap.Push(minQ, num)
			return
		}
		heap.Push(minQ, heap.Pop(maxQ).(int))
		heap.Push(maxQ, -num)

	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHp.Len() == 0 {
		return 0
	}
	maxQ, minQ := &this.maxHp, &this.minHp

	if maxQ.Len() == minQ.Len() {
		return float64(-heap.Pop(maxQ).(int)+heap.Pop(minQ).(int)) / 2
	} else {
		return float64(-heap.Pop(maxQ).(int))
	}
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
