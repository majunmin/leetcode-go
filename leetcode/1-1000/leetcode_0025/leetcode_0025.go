// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-26 00:37
package leetcode_0025

import . "github.com/majunmin/leetcode-go/common"

//https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}
	//  递归, 仅 reverse 前 K 个 元素， next, 指向下一个带翻转的链表的第一个节点
	next := head
	for i := 0; i < k; i++ {
		if next == nil {
			// 不足k 个,返回原链表
			return head
		}
		next = next.Next
	}

	//反转当前链表
	newHead := reverseList(head, next)
	// 反转后面的链表
	head.Next = reverseKGroup(next, k)
	return newHead
}

// 左闭右开区间
func reverseList(head *ListNode, tail *ListNode) *ListNode {
	var pre *ListNode
	for head != tail {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
