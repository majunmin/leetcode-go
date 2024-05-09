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

// 二分查找法, 查找第k小的数
// 题解: https://leetcode.cn/problems/median-of-two-sorted-arrays/solutions/210764/di-k-xiao-shu-jie-fa-ni-zhen-de-dong-ma-by-geek-8m/?envType=study-plan-v2&envId=top-100-liked
func solution2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	l, r := (m+n+1)>>1, (m+n+2)>>1
	return float64(getKth(nums1, 0, m-1, nums2, 0, n-1, l)+getKth(nums1, 0, m-1, nums2, 0, n-1, r)) / 2.0
}

func getKth(nums1 []int, s1 int, e1 int, nums2 []int, s2 int, e2 int, k int) int {
	len1, len2 := e1-s1+1, e2-s2+1
	// 保证  len(nums1) < len(nums2)
	if len1 > len2 {
		return getKth(nums2, s2, e2, nums1, s1, e1, k)
	}
	if len1 == 0 {
		return nums2[s2+k-1]
	}
	// len1 > 0 && len2 > 0
	if k == 1 {
		return min(nums1[s1], nums2[s2])
	}
	i, j := s1+min(len1, k/2)-1, s2+min(len2, k/2)-1
	if nums1[i] > nums2[j] {
		return getKth(nums1, s1, e1, nums2, j+1, e2, k-(j-s2+1))
	}
	return getKth(nums1, i+1, e1, nums2, s2, e2, k-(j-s1+1))
}
