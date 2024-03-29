package leetcode_0141

import . "github.com/majunmin/leetcode-go/common"

func hasCycle(head *ListNode) bool {
	// param check
	if head == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
