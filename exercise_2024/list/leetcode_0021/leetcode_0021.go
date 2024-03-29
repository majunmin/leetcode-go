package leetcode_0021

import . "github.com/majunmin/leetcode-go/common"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	return iterSolution(list1, list2)
}

func iterSolution(list1 *ListNode, list2 *ListNode) *ListNode {
	// param check
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	//
	dummyNode := &ListNode{}
	pre := dummyNode
	for list1 != nil && list2 != nil {
		var node *ListNode
		if list1.Val < list2.Val {
			node = list1
			list1 = list1.Next
		} else {
			node = list2
			list2 = list2.Next
		}
		pre.Next = node
		pre = pre.Next
	}
	if list1 != nil {
		pre.Next = list1
	}
	if list2 != nil {
		pre.Next = list2
	}
	return dummyNode.Next
}

func recursiveSolution(list1 *ListNode, list2 *ListNode) *ListNode {
	// param check
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	//
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
