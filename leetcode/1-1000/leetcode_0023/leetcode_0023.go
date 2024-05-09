package leetcode_0023

import . "github.com/majunmin/leetcode-go/common"

func mergeKLists(lists []*ListNode) *ListNode {
	// param check
	if len(lists) == 0 {
		return nil
	}
	//自底向上的  归并排序
	for len(lists) > 1 {
		left, right := 0, len(lists)-1
		for left < right {
			l1 := merge(lists[left], lists[right])
			lists[left] = l1
			lists = lists[:len(lists)-1]
			left++
			right--
		}
	}
	return lists[0]
}

// 有序链表的合并最有效的方法就是   归并排序
func solution1(lists []*ListNode) *ListNode {
	// param check
	left, right := 0, len(lists)-1
	return mergeKListsHelper(lists, left, right)
}

func mergeKListsHelper(lists []*ListNode, left int, right int) *ListNode {
	if left == right {
		return lists[0]
	}
	mid := left + (right-left)>>1
	l1 := mergeKListsHelper(lists, left, mid)
	l2 := mergeKListsHelper(lists, mid+1, right)
	return merge(l1, l2)
}

// 合并两个有序链表
func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	prev := dummyNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	//
	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}
	return dummyNode.Next
}
