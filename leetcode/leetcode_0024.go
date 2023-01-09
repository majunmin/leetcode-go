package leetcode

import . "github.com/majunmin/leetcode-go/common"

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyNode := &ListNode{}
	dummyNode.Next = head

	// processor
	pre, curNode := dummyNode, head

	for curNode != nil && curNode.Next != nil {
		pre.Next = curNode.Next
		curNode.Next = pre.Next.Next
		pre.Next.Next = curNode

		pre = curNode
		curNode = pre.Next
	}
	return dummyNode.Next

}
