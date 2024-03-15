package leetcode_0148

import (
	"container/heap"

	. "github.com/majunmin/leetcode-go/common"
)

// https://leetcode.cn/problems/sort-list/?envType=study-plan-v2&envId=top-100-liked
func sortList(head *ListNode) *ListNode {
	// 1. 常规解法， 往一个 dummyNode里面插值
	// 2. 利用 大顶堆
	return heapSort(head)
}

type maxHeap []*ListNode

func newMaxHeap() *maxHeap {
	hp := new(maxHeap)
	heap.Init(hp)
	return hp
}

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i, j int) bool {
	return m[i].Val < m[j].Val
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(*ListNode))
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

func heapSort(head *ListNode) *ListNode {
	hp := newMaxHeap()
	for head != nil {
		hp.push(head)
		head = head.Next
	}
	dummyNode := &ListNode{}
	for hp.Len() > 0 {
		item := hp.pop().(*ListNode)
		item.Next = dummyNode.Next
		dummyNode.Next = item
	}
	return dummyNode.Next
}

// 超时 😂
func insertSort(head *ListNode) *ListNode {
	dummyNode := &ListNode{}
	prev := dummyNode
	for head != nil {
		tmp := head
		head = head.Next
		// 将tmp 节点插入  dummyList 有序链表里面
		for prev.Next != nil {
			if prev.Next.Val >= tmp.Val {
				break
			}
			prev = prev.Next
		}
		tmp.Next = prev.Next
		prev.Next = tmp
		// 还原 perv 节点
		prev = dummyNode
	}
	return dummyNode.Next
}
