package leetcode_0086

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/partition-list/?envType=study-plan-v2&envId=top-interview-150
func partition(head *ListNode, x int) *ListNode {
	var (
		smallHead = &ListNode{}
		smallList = smallHead
		largeHead = &ListNode{}
		largeList = largeHead
	)
	cur := head
	for cur != nil {
		if cur.Val < x {
			smallList.Next = cur
			smallList = smallList.Next
		} else {
			largeList.Next = cur
			largeList = largeList.Next
		}
		cur = cur.Next
	}
	smallList.Next = largeHead.Next
	// 防止成环
	largeList.Next = nil
	return smallHead.Next
}
