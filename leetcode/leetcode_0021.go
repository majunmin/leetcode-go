package leetcode

import . "github.com/majunmin/leetcode-go/common"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummyNode ListNode
	pre := dummyNode
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = list1
			list1 = list1.Next
			continue
		}
		pre.Next = list2
		list2 = list2.Next
	}

	return dummyNode.Next
}
