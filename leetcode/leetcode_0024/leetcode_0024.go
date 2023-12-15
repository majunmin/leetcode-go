// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-21 20:39
package leetcode_0024

import "github.com/majunmin/leetcode-go/common"

// swapPairs Definition for singly-linked link_list.
func swapPairs(head *common.ListNode) *common.ListNode {
	return iterSolution(head)

}

func iterSolution(head *common.ListNode) *common.ListNode {
	// terminate
	if head == nil || head.Next == nil {
		return head
	}
	// process
	newHead := head.Next
	head.Next = swapPairs(head.Next.Next)
	newHead.Next = head
	return newHead
}

func recursiveSolution(head *common.ListNode) *common.ListNode {
	dummyNode := &common.ListNode{}
	dummyNode.Next = head

	prev := dummyNode
	for prev.Next != nil && prev.Next.Next != nil {
		aNode := prev.Next
		prev.Next = aNode.Next
		aNode.Next = aNode.Next.Next
		prev.Next.Next = aNode
		// move prev pointer
		prev = aNode
	}
	return dummyNode.Next
}
