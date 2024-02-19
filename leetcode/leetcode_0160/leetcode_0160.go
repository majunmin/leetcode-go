package leetcode_0160

import . "github.com/majunmin/leetcode-go/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//https://leetcode.cn/problems/intersection-of-two-linked-lists/?envType=study-plan-v2&envId=top-100-liked
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 快慢指针: 判断链表是否存在环
	//
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	// 如果不相交, headA == nil
	return pA
}
