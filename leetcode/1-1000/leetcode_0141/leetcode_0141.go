// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-22 22:42
package leetcode_0141

import . "github.com/majunmin/leetcode-go/common"

func hasCycle(head *ListNode) bool {
	// 判断链表是否存在环, 快慢指针解决
	fast := head
	slow := head
	for fast != nil && fast.Next != nil && slow != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
