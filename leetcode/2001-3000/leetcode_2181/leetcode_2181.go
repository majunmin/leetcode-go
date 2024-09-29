package leetcode_2181

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/merge-nodes-in-between-zeros/
func mergeNodes(head *ListNode) *ListNode {
	return solution1(head)
}

// 尝试一下在原地修改

// 新建一个链表.
func solution1(head *ListNode) *ListNode {
	// 在原地修改.
	if head == nil {
		return nil
	}
	var (
		cur       = head.Next
		dummyNode = new(ListNode)
		curVal    int
	)
	prev := dummyNode
	for ; cur != nil; cur = cur.Next {
		if cur.Val == 0 {
			prev.Next = &ListNode{Val: curVal}
			prev = prev.Next
			curVal = 0
			continue
		}
		curVal += cur.Val
	}
	return dummyNode.Next
}
