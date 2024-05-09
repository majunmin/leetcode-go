package leetcode_0019

import (
	. "github.com/majunmin/leetcode-go/common"
)

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/?envType=study-plan-v2&envId=top-100-liked
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//
	dummyNode := &ListNode{}
	prev := dummyNode
	prev.Next = head
	p := head
	for i := 0; i <= n; i++ {
		if p == nil {
			return head
		}
		p = p.Next
	}
	for p != nil {
		prev = prev.Next
		p = p.Next
	}
	prev.Next = prev.Next.Next
	return dummyNode.Next
}
