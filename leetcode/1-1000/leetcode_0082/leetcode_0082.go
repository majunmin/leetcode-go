package leetcode_0082

import . "github.com/majunmin/leetcode-go/common"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var (
		dummyNode = &ListNode{Val: 300}
		prev      = dummyNode
	)
	dummyNode.Next = head
	for prev.Next != nil && prev.Next.Next != nil {
		if prev.Next.Val == prev.Next.Next.Val {
			// del
			delVal := prev.Next.Val
			for prev.Next != nil && prev.Next.Val == delVal {
				prev.Next = prev.Next.Next
			}
			continue
		}
		prev = prev.Next
	}
	return dummyNode.Next
}

func solution2(head *ListNode) *ListNode {
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
		cur = pre.Next
	}

	return dummyNode.Next
}
