package lcr_141

import . "github.com/majunmin/leetcode-go/common"

//https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func trainningPlan(head *ListNode) *ListNode {
	// reverse list
	if head == nil || head.Next == nil {
		return head
	}
	next := trainningPlan(head.Next)
	// reverse
	head.Next.Next = head
	head.Next = nil
	return next
}

func trainningPlanIter(head *ListNode) *ListNode {
	dummyNode := new(ListNode)
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = dummyNode.Next
		dummyNode.Next = cur
		cur = tmp
	}
	return dummyNode.Next
}
