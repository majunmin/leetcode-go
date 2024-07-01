package lcr_078

import . "github.com/majunmin/leetcode-go/common"

// 1. 小顶堆
// 2. 归并排序
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//https://leetcode.cn/problems/vvXgSW/
func mergeKLists(lists []*ListNode) *ListNode {
	return mergeSort(lists, 0, len(lists)-1)
}

func mergeSort(lists []*ListNode, left int, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)/2
	list1 := mergeSort(lists, left, mid)
	list2 := mergeSort(lists, mid+1, right)
	return mergeTwoList(list1, list2)
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoList(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoList(list1, list2.Next)
		return list2
	}
}
