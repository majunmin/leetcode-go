package leetcode_0004

import (
	"container/heap"
	"math"
)

type maxHeap []int

func newMaxHeap() *maxHeap {
	maxHp := new(maxHeap)
	heap.Init(maxHp)
	return maxHp
}

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

func (m *maxHeap) push(x any) {
	heap.Push(m, x)
}

func (m *maxHeap) pop() any {
	return heap.Pop(m)
}

func (m *maxHeap) top() any {
	return (*m)[0]
}

// https://leetcode.cn/problems/median-of-two-sorted-arrays/?envType=study-plan-v2&envId=top-100-liked
// 堆解法 -1
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	m, n := len(nums1), len(nums2)
	if m+n == 0 {
		return float64(0)
	}
	var (
		mHeap = newMaxHeap()

		targetNum = (m + n + 2) >> 1
		l, r      int
	)

	for i := 0; i < targetNum; i++ {
		lval, rval := math.MaxInt, math.MaxInt
		if l < m {
			lval = nums1[l]
		}
		if r < n {
			rval = nums2[r]
		}
		if lval < rval {
			l++
		} else {
			r++
		}
		mHeap.push(min(lval, rval))
	}
	if (m+n)&1 == 1 {
		return float64(mHeap.pop().(int))
	}
	return (float64(mHeap.pop().(int)) + float64(mHeap.pop().(int))) / 2.0
}
