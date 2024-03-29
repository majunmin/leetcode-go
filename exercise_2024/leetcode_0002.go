package exercise_2024

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/add-two-numbers/?envType=study-plan-v2&envId=top-100-liked
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := new(ListNode)
	prev := dummyNode
	p := 0 // 进位
	for l1 != nil || l2 != nil {
		cur := p
		if l1 != nil {
			cur += l1.Val
		}
		if l2 != nil {
			cur += l2.Val
		}
		prev.Next = &ListNode{Val: cur % 10}
		p = cur / 10
		prev = prev.Next
	}
	if p > 0 {
		prev.Next = &ListNode{Val: 1}
	}
	return dummyNode.Next
}
