package offer_064

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/?envType=study-plan&id=lcof
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	//return recursiveSolution(head)
	return iterSolution(head)
}

func iterSolution(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//
	dummyNode := &ListNode{}
	node := head
	for node != nil {
		tmp := node.Next
		node.Next = dummyNode.Next
		dummyNode.Next = node
		node = tmp
	}
	return dummyNode.Next
}

func recursiveSolution(head *ListNode) *ListNode {
	// terminate
	if head == nil || head.Next == nil {
		return head
	}
	// process

	next := recursiveSolution(head.Next)
	head.Next.Next = head
	head.Next = nil
	return next
}
