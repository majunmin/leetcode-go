package leetcode_0082

import . "github.com/majunmin/leetcode-go/common"

func deleteDuplicates(head *ListNode) *ListNode {
	dummyNode := ListNode{Val: 101}
	dummyNode.Next = head
	cur := head
	pre := &dummyNode
	for cur != nil && cur.Next != nil {
		if cur.Val != cur.Next.Val {
			pre = pre.Next
			cur = cur.Next
			continue
		}
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		pre.Next = cur.Next
		cur.Next = nil
		pre = pre.Next
		cur = pre.Next
	}

	return dummyNode.Next
}
