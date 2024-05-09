package leetcode_0083

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == cur.Val {
			// remove cur.Next
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return head
}
