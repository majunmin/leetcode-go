package leetcode_0004

import (
	"container/heap"
	"sort"
)

type maxHp struct {
	sort.IntSlice
}

func newMaxHp() *maxHp {
	return &maxHp{}
}

func (m *maxHp) Less(i, j int) bool {
	return !m.IntSlice.Less(i, j)
}

func (m *maxHp) Push(x any) {
	m.IntSlice = append(m.IntSlice, x.(int))
}

func (m *maxHp) Pop() any {
	item := m.IntSlice[len(m.IntSlice)-1]
	m.IntSlice = m.IntSlice[:len(m.IntSlice)-1]
	return item
}

// https://leetcode.cn/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// param check
	llen, rlen := len(nums1), len(nums2)
	idx := (llen + rlen) >> 1
	var idx1, idx2 int
	maxHp := newMaxHp()
	heap.Init(maxHp)
	var cnt int
	for idx1 < len(nums1) || idx2 < len(nums2) {
		if cnt == idx+1 {
			break
		}
		if idx1 < len(nums1) && idx2 < len(nums2) {
			if nums1[idx1] < nums2[idx2] {
				heap.Push(maxHp, nums1[idx1])
				idx1++
			} else {
				heap.Push(maxHp, nums2[idx2])
				idx2++
			}
		} else if idx1 < len(nums1) {
			heap.Push(maxHp, nums1[idx1])
			idx1++
		} else if idx2 < len(nums2) {
			heap.Push(maxHp, nums2[idx2])
			idx2++
		}
		cnt++
	}
	total := len(nums1) + len(nums2)
	val1 := heap.Pop(maxHp).(int)
	if total&1 == 1 {
		return float64(val1)
	} else {
		val2 := heap.Pop(maxHp).(int)
		return float64(val1+val2) / float64(2.0)
	}
}
