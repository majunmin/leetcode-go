package leetcode_0141

import . "github.com/majunmin/leetcode-go/common"

func hasCycle(head *ListNode) bool {
	// param check
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if slow == fast {
			return true
		}
	}
	return false
}
