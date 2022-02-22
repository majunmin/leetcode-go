// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-22 23:15
package leetcode_0092

import . "github.com/majunmin/leetcode-go/common"

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 1. 先找到 leftNode.prev
	// 2. 后面执行两个操作:
	//    - delete leftNode.Next
	//    - insert prev.Next

	prev := &ListNode{Next: head}
	for i := 1; i < left; i++ {
		prev = prev.Next
	}
	leftNode := prev.Next

	// 操作 right-left 次
	for i := 1; i <= (right - left); i++ {
		nextNode := leftNode.Next
		// del
		leftNode.Next = leftNode.Next.Next
		// insert
		nextNode.Next = prev.Next
		prev.Next = nextNode
	}
	if left == 1 {
		return prev.Next
	}
	return head
}
