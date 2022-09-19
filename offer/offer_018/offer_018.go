package offer_018

import . "github.com/majunmin/leetcode-go/offer"

func deleteNode(head *ListNode, val int) *ListNode {
	// check parameter
	if head == nil {
		return nil
	}
	dummyNode := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := dummyNode

	for pre != nil && pre.Next != nil {
		// del  node
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		}
		pre = pre.Next
	}

	return dummyNode.Next
}
