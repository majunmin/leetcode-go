package lcr_025

import . "github.com/majunmin/leetcode-go/common"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 利用栈
	var (
		stack1 = make([]int, 0)
		stack2 = make([]int, 0)
	)
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	dummyNode := new(ListNode)
	p := 0 // 进位
	for len(stack1) > 0 || len(stack2) > 0 {
		//pop stack1
		if len(stack1) > 0 {
			p += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		//pop stack2
		if len(stack2) > 0 {
			p += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		// 头插
		//prev.Next = &ListNode{Val: p % 10}
		newNode := &ListNode{Val: p % 10}
		newNode.Next = dummyNode.Next
		dummyNode.Next = newNode
		p /= 10
	}
	if p == 1 {
		newNode := &ListNode{Val: p}
		newNode.Next = dummyNode.Next
		dummyNode.Next = newNode
	}
	// GC
	return dummyNode.Next
}
