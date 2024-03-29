package exercise_2024

import . "github.com/majunmin/leetcode-go/common"

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	// 归并排序
	l, r := 0, len(lists)-1
	return mergeKListsHelper(lists, l, r)
}

func mergeKListsHelper(lists []*ListNode, l int, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	mid := l + (r-l)>>1
	l1 := mergeKListsHelper(lists, l, mid)
	l2 := mergeKListsHelper(lists, mid+1, r)
	return merge(l1, l2)
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	}
	l2.Next = merge(l1, l2.Next)
	return l2
}
