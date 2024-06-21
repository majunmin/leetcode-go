package leetcode_0445

import . "github.com/majunmin/leetcode-go/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		stack1    = make([]int, 0)
		stack2    = make([]int, 0)
		dummyNode = new(ListNode)
		p         int // 进位
	)
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	for len(stack1) > 0 || len(stack2) > 0 {
		cur := p
		if len(stack1) > 0 {
			cur += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			cur += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}

		node := &ListNode{Val: cur % 10}
		node.Next = dummyNode.Next
		dummyNode.Next = node
		p = cur / 10
	}
	if p == 1 {
		node := &ListNode{Val: 1}
		node.Next = dummyNode.Next
		dummyNode.Next = node
	}
	return dummyNode.Next
}
