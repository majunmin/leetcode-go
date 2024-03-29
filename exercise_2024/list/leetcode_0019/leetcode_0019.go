package leetcode_0019

import (
	. "github.com/majunmin/leetcode-go/common"
)

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/?envType=study-plan-v2&envId=top-100-liked
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummyNode := &ListNode{Next: head}
	prev, tail := dummyNode, dummyNode
	var i int
	for i <= n && tail != nil {
		tail = tail.Next
		i++
	}

	// 想想这里为什么是 n+1
	if tail == nil && i == n+1 {
		return head.Next
	}
	for tail != nil {
		tail = tail.Next
		prev = prev.Next
	}
	// delete prev.Next
	prev.Next = prev.Next.Next
	return head
}
