package lcr_171

import . "github.com/majunmin/leetcode-go/common"

//https://leetcode.cn/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 返回两个链表的交点
	n1, n2 := headA, headB
	for n1 != n2 {
		if n1 == nil {
			n1 = n2
		} else {
			n1 = n1.Next
		}
		if n2 == nil {
			n2 = n1
		} else {
			n2 = n2.Next
		}
	}
	return n1
}
