package leetcode_0147

import . "github.com/majunmin/leetcode-go/common"

func insertionSortList(head *ListNode) *ListNode {
	// check
	if head == nil || head.Next == nil {
		return head
	}
	dummyNode := &ListNode{}
	cur := head
	for cur != nil {
		pre := dummyNode
		tmp := cur.Next
		cur.Next = nil
		for pre.Next != nil {
			if cur.Val <= pre.Next.Val {
				break
			}
			pre = pre.Next
		}
		cur.Next = pre.Next
		pre.Next = cur
		cur = tmp
	}
	return dummyNode.Next
}
