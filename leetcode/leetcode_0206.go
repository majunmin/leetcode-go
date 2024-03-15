package leetcode

import . "github.com/majunmin/leetcode-go/common"

func reverseList(head *ListNode) *ListNode {
	return solution3(head)
}

func solution3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		cur.Next = prev
		prev = cur
		cur = cur.Next
	}
	return prev
}

func iterSolution(head *ListNode) *ListNode {
	// if
	if head == nil || head.Next == nil {
		return head
	}
	var dummyNode ListNode
	for head != nil {
		tmp := head.Next
		head.Next = dummyNode.Next
		dummyNode.Next = head
		head = tmp
	}
	return dummyNode.Next

}

func recursiveSolution(head *ListNode) *ListNode {
	// param check
	if head == nil || head.Next == nil {
		return head
	}
	//
	tail := recursiveSolution(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tail
}
