package lcr_142

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func trainningPlan(l1 *ListNode, l2 *ListNode) *ListNode {
	// 合并两个有序链表
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = trainningPlan(l1.Next, l2)
		return l1
	}
	l2.Next = trainningPlan(l1, l2.Next)
	return l2
}
