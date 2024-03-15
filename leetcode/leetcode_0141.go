package leetcode

import . "github.com/majunmin/leetcode-go/common"

func hasCycle(head *ListNode) bool {
	// param check
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
