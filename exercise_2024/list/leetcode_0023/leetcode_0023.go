package leetcode_0023

import (
	"container/heap"

	. "github.com/majunmin/leetcode-go/common"
)

func mergeKLists(lists []*ListNode) *ListNode {
	// param check
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	// merge sort
	return heapSolution(lists)

}

func mergeSortSolution(lists []*ListNode) *ListNode {
	left, right := 0, len(lists)-1
	return mergeSortLists(lists, left, right)
}

// 2. 堆的解法
type minHeap struct {
	items []*ListNode
}

func newMinHeap() *minHeap {
	return &minHeap{items: make([]*ListNode, 0)}
}

func (m *minHeap) Len() int {
	return len(m.items)
}

func (m *minHeap) Less(i, j int) bool {
	return m.items[i].Val < m.items[j].Val
}

func (m *minHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

func (m *minHeap) Push(x any) {
	m.items = append(m.items, x.(*ListNode))
}

func (m *minHeap) Pop() any {
	item := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return item
}

func heapSolution(lists []*ListNode) *ListNode {
	mHp := newMinHeap()
	heap.Init(mHp)
	dummyNode := &ListNode{}
	pre := dummyNode
	for _, node := range lists {
		heap.Push(mHp, node)
	}
	for len(mHp.items) > 0 {
		item := heap.Pop(mHp).(*ListNode)
		tmp := item.Next
		pre.Next = item
		pre = pre.Next
		pre.Next = nil
		if tmp != nil {
			heap.Push(mHp, tmp)
		}
	}
	return dummyNode.Next
}

// 1. 分治解法
func mergeSortLists(lists []*ListNode, left int, right int) *ListNode {
	if left > right {
		return nil
	}
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)>>1
	l1, l2 := mergeSortLists(lists, left, mid), mergeSortLists(lists, mid+1, right)
	return mergeLists(l1, l2)
}

func mergeLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	//
	dummyNode := &ListNode{}
	pre := dummyNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}
	return dummyNode.Next
}
