package leetcode_0234

import (
	. "github.com/majunmin/leetcode-go/common"
)

func isPalindrome(head *ListNode) bool {
	// param check
	if head == nil || head.Next == nil {
		return true
	}
	//
	slow, fast := head, head
	for fast != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// reverse 后半段
	next := reverseList(slow.Next)
	l1, l2 := head, next
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
