package leetcode_0206

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/reverse-linked-list/?envType=study-plan-v2&envId=top-100-liked
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return next
}
