package leetcode_0002

import (
	. "github.com/majunmin/leetcode-go/common"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return iterSolution(l1, l2)
}

// 可以优化的一点是: 将
func iterSolution(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	prev := dummyNode
	var p int
	for l1 != nil && l2 != nil && p != 0 {
		curVal := p
		if l1 != nil {
			curVal += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			curVal += l2.Val
			l2 = l2.Next
		}
		p = curVal / 10
		node := &ListNode{Val: curVal % 10}
		prev.Next = node
		prev = prev.Next
	}
	return dummyNode.Next
}

func recursiveSolution(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersInner(l1, l2, 0)
}

func addTwoNumbersInner(l1 *ListNode, l2 *ListNode, p int) *ListNode {
	// terminate
	if l1 == nil && l2 == nil {
		if p == 0 {
			return nil
		}
		return &ListNode{Val: p}
	}
	//
	var curVal int
	if l1 != nil {
		curVal += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		curVal += l2.Val
		l2 = l2.Next
	}
	curVal += p
	p = curVal / 10
	curVal = curVal % 10

	node := &ListNode{Val: curVal}
	node.Next = addTwoNumbersInner(l1, l2, p)
	return node
}
