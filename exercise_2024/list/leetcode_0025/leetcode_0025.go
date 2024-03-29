package leetcode_0025

import . "github.com/majunmin/leetcode-go/common"

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{Next: head}
	pre, tail := dummyNode, dummyNode
	next := pre.Next
	for tail != nil {
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				break
			}
		}
		tmp := tail.Next
		tail.Next = nil
		reverse(pre.Next)
		pre.Next = tail
		next.Next = tmp
		pre, tail = next, next
		next = pre.Next
	}
	return dummyNode.Next
}

func reverKlist1(head *ListNode, k int) *ListNode {
	// param check
	if k < 1 {
		panic("invalid param")
	}
	if k == 1 {
		return head
	}
	// process
	return reversePreKList(head, k)
}

// reverse 前k个节点
func reversePreKList(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for i := 1; i < k; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return head
	}
	next := cur.Next
	cur.Next = nil
	newHead := reverse(head)
	head.Next = reversePreKList(next, k)
	return newHead //
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
