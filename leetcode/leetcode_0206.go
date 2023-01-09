package leetcode

import . "github.com/majunmin/leetcode-go/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyNode := &ListNode{}
	current := head
	for current != nil {
		tmp := current.Next
		// insert node into dummyList
		current.Next = dummyNode.Next
		dummyNode.Next = current
		current = tmp
	}

	return dummyNode.Next
}
