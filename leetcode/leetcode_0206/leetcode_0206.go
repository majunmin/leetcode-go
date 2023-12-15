// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-24 23:52
package leetcode_0206

import "github.com/majunmin/leetcode-go/common"

/**
 * Definition for singly-linked link_list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *common.ListNode) *common.ListNode {
	return recursiveSolution(head)
}

func recursiveSolution(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := recursiveSolution(head.Next)
	head.Next.Next = head
	next.Next = nil // 断尾
	return next
}

func iterSolution(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}
	var dummyNode common.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = dummyNode.Next
		dummyNode.Next = cur
		cur = next
	}
	return dummyNode.Next
}
